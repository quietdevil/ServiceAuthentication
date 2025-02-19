package authentication

import (
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
	"github.com/quietdevil/ServiceAuthentication/internal/service"
)

type AuthenticationService struct {
	reposAccess repository.AuthenticationRepository
}

func NewAuthenticationService(reposAccess repository.AuthenticationRepository) service.AuthenticationService {
	return &AuthenticationService{reposAccess: reposAccess}
}
