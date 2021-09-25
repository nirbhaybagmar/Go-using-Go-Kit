package schema

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CreateUserRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CreateUserResponse struct {
	Ok string `json:"ok"`
}

type GetUserRequest struct {
	Id string `json:"id,omitempty"`
}

type GetUserResponse struct {
	Email string `json:"email,omitempty"`
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	req := GetUserRequest{
		Id: vars["id"],
	}
	return req, nil

}
