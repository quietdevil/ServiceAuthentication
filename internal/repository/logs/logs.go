package logs

import (
	"context"
	"serviceauth/internal/repository"
	l "serviceauth/internal/repository/logs/model"

	db "github.com/quietdevil/Platform_common/pkg/db"
)

type Logs struct {
	DB db.Client
}

func NewLogs(db db.Client) repository.Logger {
	return &Logs{DB: db}
}

func (l *Logs) Create(ctx context.Context, log l.Logs) error {
	q := db.Query{
		Name:     log.Name,
		QueryStr: "INSERT INTO logs (name, description) VALUES ($1, $2)",
	}

	_, err := l.DB.DB().ContextExec(ctx, q, log.Name, log.Description)
	if err != nil {
		return err
	}
	return nil
}
