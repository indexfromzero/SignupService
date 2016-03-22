package contract

import "github.com/microbusinesses/SignupService/domain"

type SignupService interface {

	//Signup a new user
	Signup(user domain.User) error
}
