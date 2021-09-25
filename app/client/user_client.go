package client

import (
	"context"
	"gokit/app/user"
)

func (c client) CreateUser(ctx context.Context, user user.User) error {
	sql := `INSERT INTO users (id, email, password) values($1, $2, $3) `

	if user.Email == "" || user.Password == "" {
		return ClientError
	}
	_, err := c.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (c client) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := c.db.QueryRow("SELECT email from users where id=$1", id).Scan(&email)
	if err != nil {
		return "", ClientError
	}
	return email, nil
}
