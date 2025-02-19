package access

import (
	"context"
	"github.com/quietdevil/Platform_common/pkg/db"
	"github.com/quietdevil/ServiceAuthentication/internal/repository"
)

type accessRepos struct {
	DB db.Client
}

func NewAccessRepository(DB db.Client) repository.AccessRepository {
	return &accessRepos{DB: DB}
}

func (a *accessRepos) Role(ctx context.Context, endpoint string) (string, error) {
	err := a.DB.DB().Ping(ctx)
	if err != nil {
		return "", err
	}

	q := db.Query{
		QueryStr: "SELECT role FROM endpoints WHERE endpoint=$1",
	}

	row, err := a.DB.DB().ContextQuery(ctx, q, endpoint)
	if err != nil {
		return "", err
	}
	var role string

	for row.Next() {
		if err := row.Scan(&role); err != nil {
			return "", err
		}
	}
	return role, nil

}
