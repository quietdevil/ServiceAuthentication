package access

import (
	"github.com/quietdevil/ServiceAuthentication/internal/config"
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
	"github.com/quietdevil/ServiceAuthentication/internal/service"
)

type AccessService struct {
	reposAccess repository.AccessRepository
	configAuth  *config.AuthenticationConfig
}

func NewAccessService(reposAccess repository.AccessRepository, conf *config.AuthenticationConfig) service.AccessService {
	return &AccessService{reposAccess: reposAccess, configAuth: conf}
}
