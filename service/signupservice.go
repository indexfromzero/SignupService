package service

import "github.com/microbusinesses/Micro-Businesses-Core/system"

type SignupService struct{}

func (SignupService) Create(object interface{}) error {
	return nil
}

func (SignupService) Get(id system.UUID) interface{} {
	return nil
}

func (SignupService) Update(id system.UUID, object interface{}) error {
	return nil
}

func (SignupService) Delete(id system.UUID) error {
	return nil
}
