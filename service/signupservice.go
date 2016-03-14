package service

import "github.com/microbusinesses/Micro-Businesses-Core/system"

type SignupService struct{}

func (SignupService) create(object interface{}) error {
	return nil
}

func (SignupService) get(id system.UUID) interface{} {
	return nil
}

func (SignupService) update(id system.UUID, object interface{}) error {
	return nil
}

func (SignupService) delete(id system.UUID) error {
	return nil
}
