package authentication

import (
	"github.com/quietdevil/ServiceAuthentication/internal/config"
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
	"github.com/quietdevil/ServiceAuthentication/internal/service"
)

type AuthenticationService struct {
	reposAccess repository.AuthenticationRepository
	configAuth  *config.AuthenticationConfig
}

func NewAuthenticationService(reposAccess repository.AuthenticationRepository, conf *config.AuthenticationConfig) service.AuthenticationService {
	return &AuthenticationService{reposAccess: reposAccess, configAuth: conf}
}
