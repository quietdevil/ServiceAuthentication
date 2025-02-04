package user

import (
	"serviceauth/internal/service"
	desc "serviceauth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthenticationServer
	userService service.UserService
}

func NewImplementation(service service.UserService) *Implementation {
	return &Implementation{
		userService: service,
	}
}
