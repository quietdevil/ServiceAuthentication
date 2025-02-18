package authentication

import (
	"serviceauth/internal/service"
	desc "serviceauth/pkg/auth_v1"
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
