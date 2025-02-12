package repository

import (
	"context"
	"serviceauth/internal/model"
	l "serviceauth/internal/repository/logs/model"
)

// urlExample := "postgres://username:password@localhost:5432/database_name"
// var database = fmt.Sprintf("postgres://%v:%v@localhost:%v/%v", os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_PORT"), os.Getenv("PG_DATABASE_NAME"))
type UserRepository interface {
	Get(context.Context, int) (*model.User, error)
	Create(context.Context, *model.UserInfo) (int, error)
	Delete(context.Context, int) error
	Update(context.Context, *model.UserUpdate) error
}

type Logger interface {
	Create(context.Context, l.Logs) error
}
