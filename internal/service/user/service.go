package user

import (
	"serviceauth/internal/repository"
	"serviceauth/internal/service"

	db "github.com/quietdevil/Platform_common/pkg/db"
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
