package access

import (
	"serviceauth/internal/repository"
	"serviceauth/internal/service"
)

type AccessService struct {
	reposAccess repository.AccessRepository
}

func NewAccessService(reposAccess repository.AccessRepository) service.AccessService {
	return &AccessService{reposAccess: reposAccess}
}
