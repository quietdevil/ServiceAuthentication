package user

import (
	"github.com/quietdevil/ServiceAuthentication/internal/service"
	desc "github.com/quietdevil/ServiceAuthentication/pkg/auth_user_v1"
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
