package user

import (
	"serviceauth/internal/client/db"
	"serviceauth/internal/repository"
	"serviceauth/internal/service"
)

type serviceUser struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
	logs           repository.Logger
}

func NewService(repos repository.UserRepository, tx db.TxManager, logs repository.Logger) service.UserService {
	return &serviceUser{
		userRepository: repos,
		txManager:      tx,
		logs:           logs,
	}
}
