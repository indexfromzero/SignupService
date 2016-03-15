package main

import (
	"github.com/microbusinesses/Micro-Businesses-Core/system"
	"github.com/microbusinesses/SignupService/domain"
	"github.com/microbusinesses/SignupService/service"
	"golang.org/x/net/context"
)

type createUserRequest struct {
	User domain.User `json:"user"`
}

type createUserResponse struct {
	UserId system.UUID `json:"userid"`
	Err    string      `json:"err,omitempty"`
}

type getUserRequest struct {
	UserId system.UUID `json:"userid"`
}

type getUserResponse struct {
	User domain.User `json:"user"`
}

type updateUserRequest struct {
	User domain.User `json:"user"`
}

type updateUserResonse struct {
	Err string `json:"err,omitempty"`
}

type deleteUserRequest struct {
	UserId system.UUID `json:"userid"`
}

type deleteUserResponse struct {
	Err string `json:"err,omitempty"`
}

func makeCreateUserEndpoint(svc service.SignupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		userId, err := svc.Create(req.User)
		if err != nil {
			return createUserResponse{userId, err.Error()}, nil
		}
		return createUserResponse{userId, ""}, nil
	}
}
