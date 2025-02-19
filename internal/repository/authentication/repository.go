package authentication

import (
	"github.com/quietdevil/Platform_common/pkg/db"
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
)

type AuthRepos struct {
	DB db.Client
}

func NewAuthRepos(DB db.Client) repository.AuthenticationRepository {
	return &AuthRepos{DB: DB}
}
