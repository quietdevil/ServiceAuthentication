package user

import (
	"serviceauth/internal/service"
	desc "serviceauth/pkg/auth_user_v1"
)

type Implementation struct {
	desc.UnimplementedAuthenticationUserV1Server
	userService service.UserService
}

func NewImplementation(service service.UserService) *Implementation {
	return &Implementation{
		userService: service,
	}
}
