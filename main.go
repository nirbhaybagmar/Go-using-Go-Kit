package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gokit/app"
	"gokit/app/client"
	"gokit/app/server"
	"gokit/app/user"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var dbsource = "sqlite3:/user.db?_fk=1"

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger

	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "accounts",
			"time:", log.DefaultTimestampUTC,
			"caller:", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()
	var serv app.Service
	{
		client := client.NewClient(db, logger)
		serv = user.NewService(client, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := server.MakeEndpoints(serv)

	go func() {
		fmt.Println("Listening on port: ", *httpAddr)
		handler := server.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
