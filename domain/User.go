package domain

import "github.com/microbusinesses/Micro-Businesses-Core/system"

type User struct {
	OrganizationId system.UUID
	UserId         system.UUID
	Email          string
	Username       string
	Password       string
}
