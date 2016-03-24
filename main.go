package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/microbusinesses/SignupService/domain"
	"github.com/microbusinesses/SignupService/service"
	"golang.org/x/net/context"
)

type signupRequest struct {
	User domain.User `json:"user"`
}

type signupResponse struct {
	Err string `json:"err,omitempty"`
}

func makeSignupEndpoint(svc service.SignupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signupRequest)
		err := svc.Signup(req.User)
		if err != nil {
			return signupResponse{err.Error()}, nil
		}
		return signupResponse{""}, nil
	}
}

func main() {
	ctx := context.Background()
	svc := service.SignupService{}

	signupHandler := httptransport.NewServer(
		ctx,
		makeSignupEndpoint(svc),
		decodeSignupRequest,
		encodeSignupResponse,
	)

	http.Handle("/signup", signupHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeSignupRequest(r *http.Request) (interface{}, error) {
	var request signupRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeSignupResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
