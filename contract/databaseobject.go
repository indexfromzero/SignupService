package contract

import "github.com/microbusinesses/Micro-Businesses-Core/system"

type IDatabaseObject interface {

	//Creates a new object
	Create(object interface{}) error

	//Gets an existing object
	Get(id system.UUID) interface{}

	//Updates an existing object
	Update(id system.UUID, object interface{}) error

	//Deletes an existing object
	Delete(id system.UUID) error
}
