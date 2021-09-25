package app

import "context"

type Service interface {
	GetUser(ctx context.Context, id string) (string, error)
	CreateUser(ctx context.Context, email string, password string) (string, error)
}
