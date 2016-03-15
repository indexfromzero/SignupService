package contract

import (
	"github.com/microbusinesses/Micro-Businesses-Core/system"
	"github.com/microbusinesses/SignupService/domain"
)

type ISignupService interface {

	//Creates a new user
	Create(user domain.User) (system.UUID, error)

	//Gets an existing user
	Get(userId system.UUID) domain.User

	//Updates an existing user
	Update(user domain.User) error

	//Deletes an existing user
	Delete(id system.UUID) error
}
