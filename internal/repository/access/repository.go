package access

import (
	"context"
	"github.com/quietdevil/Platform_common/pkg/db"
	"serviceauth/internal/repository"
)

type accessRepos struct {
	DB db.Client
}

func NewAccessRepository(DB db.Client) repository.AccessRepository {
	return &accessRepos{DB: DB}
}

func (a *accessRepos) Roles(ctx context.Context, idRole int) error {
	return nil
}
