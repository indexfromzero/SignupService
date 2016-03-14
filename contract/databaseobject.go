package contract

import "github.com/microbusinesses/Micro-Businesses-Core/system"

type IDatabaseObject interface {

	//Creates a new object
	create(object interface{}) error

	//Gets an existing object
	get(id system.UUID) interface{}

	//Updates an existing object
	update(id system.UUID, object interface{}) error

	//Deletes an existing object
	delete(id system.UUID) error
}
