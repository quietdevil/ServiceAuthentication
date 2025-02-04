package user

import (
	"serviceauth/internal/repository"
	"serviceauth/internal/service"
)

type serviceUser struct {
	userRepository repository.UserRepository
}

func NewService(repos repository.UserRepository) service.UserService {
	return &serviceUser{
		userRepository: repos,
	}
}
