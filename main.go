package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
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

func main() {

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

func main() {
	ctx := context.Background()
	svc := service.SignupService{}

	createUserHandler := httptransport.NewServer(
		ctx,
		makeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	)

	http.Handle("/create-user", createUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeCreateUserRequest(r *http.Request) (interface{}, error) {
	var request createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, Err
	}
	return request, nil
}

func encodeCreateUserResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
