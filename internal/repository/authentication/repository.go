package authentication

import (
	"github.com/quietdevil/Platform_common/pkg/db"
	"serviceauth/internal/repository"
)

type AuthRepos struct {
	DB db.Client
}

func NewAuthRepos(DB db.Client) repository.AuthenticationRepository {
	return &AuthRepos{DB: DB}
}
