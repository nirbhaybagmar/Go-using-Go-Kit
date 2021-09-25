package server

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"gokit/app/schema"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(Middleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser, schema.DecodeUserReq, schema.EncodeResponse))
	r.Methods("GET").Path("/user/" +
		"{id}").Handler(httptransport.NewServer(endpoints.GetUser, schema.DecodeEmailReq, schema.EncodeResponse))

	return r
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			handler.ServeHTTP(w, r)
		})
}
