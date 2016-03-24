package service

import (
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	"github.com/microbusinesses/SignupService/domain"
)

type SignupService struct{}

func (SignupService) Signup(user domain.User) error {
	diagnostics.IsNotNilOrEmptyOrWhitespace(user.Password, "password", "User password must not be empty")

	if user.Username == "" || user.Email == "" {
		panic("Must have an email or a username")
	}

	initializeCluster()
	createKeyspace()

	if err := signupUser(user); err != nil {
		return err
	}
	return nil
}
