package server

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gokit/app"
	"gokit/app/schema"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s app.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(schema.CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return schema.CreateUserResponse{
			ok,
		}, err
	}
}

func makeGetUserEndpoint(s app.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(schema.GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return schema.GetUserResponse{
			email,
		}, err
	}
}
