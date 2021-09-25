package client

import (
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
	"gokit/app/user"
)

var ClientError = errors.New("Unable to handle Repo Request.")

type client struct {
	db     *sql.DB
	logger log.Logger
}

func NewClient(db *sql.DB, logger log.Logger) user.UserQuery {
	return &client{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}
