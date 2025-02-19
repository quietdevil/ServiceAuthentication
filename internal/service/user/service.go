package user

import (
	db "github.com/quietdevil/Platform_common/pkg/db"
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
	"github.com/quietdevil/ServiceAuthentication/internal/service"
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
