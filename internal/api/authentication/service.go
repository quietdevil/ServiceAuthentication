package authentication

import (
	"github.com/quietdevil/ServiceAuthentication/internal/service"
	desc "github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
)

type ImplementationAuthentication struct {
	desc.UnimplementedAuthenticationV1Server
	service service.AuthenticationService
}

func NewImplementationAuthentication(service service.AuthenticationService) *ImplementationAuthentication {
	return &ImplementationAuthentication{
		service: service,
	}
}
