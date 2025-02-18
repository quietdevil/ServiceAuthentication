package authentication

import (
	"serviceauth/internal/repository"
	"serviceauth/internal/service"
)

type AuthenticationService struct {
	reposAccess repository.AuthenticationRepository
}

func NewAuthenticationService(reposAccess repository.AuthenticationRepository) service.AuthenticationService {
	return &AuthenticationService{reposAccess: reposAccess}
}
