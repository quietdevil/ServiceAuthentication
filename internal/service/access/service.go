package access

import (
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
	"github.com/quietdevil/ServiceAuthentication/internal/service"
)

type AccessService struct {
	reposAccess repository.AccessRepository
}

func NewAccessService(reposAccess repository.AccessRepository) service.AccessService {
	return &AccessService{reposAccess: reposAccess}
}
