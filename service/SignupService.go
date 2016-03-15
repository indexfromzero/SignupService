package service

import (
	"github.com/microbusinesses/Micro-Businesses-Core/system"
	"github.com/microbusinesses/SignupService/domain"
)

type SignupService struct{}

func (SignupService) Create(user domain.User) (system.UUID, error) {
	return system.EmptyUUID, nil
}

func (SignupService) Get(id system.UUID) domain.User {
	return domain.User{}
}

func (SignupService) Update(user domain.User) error {
	return nil
}

func (SignupService) Delete(id system.UUID) error {
	return nil
}
