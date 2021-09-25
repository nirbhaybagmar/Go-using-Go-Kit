package user

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	"gokit/app"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserQuery interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}

type service struct {
	usrQuery UserQuery
	logger   log.Logger
}

func NewService(usrQuery UserQuery, log log.Logger) app.Service {
	return &service{usrQuery, log}
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.usrQuery.GetUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Get User", email)
	return "Success", nil
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uid, _ := uuid.NewV4()
	id := uid.String()

	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.usrQuery.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("create user", id)
	return "Success", nil
}
