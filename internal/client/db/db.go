package db

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Handler func(context.Context) error

type TxManager interface {
	ReadCommitted(context.Context, Handler) error
}

type Transactor interface {
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

type Client interface {
	DB() DB
	Close() error
}

type DB interface {
	QueryExecer
	Pinger
	Transactor
	Close()
}

type Query struct {
	Name        string
	QueryString string
}

type QueryExecer interface {
	ExecContext(context.Context, Query, ...any) (pgconn.CommandTag, error)
	QueryContext(context.Context, Query, ...any) (pgx.Rows, error)
	QueryRowContext(context.Context, Query, ...any) pgx.Row
}

type Pinger interface {
	Ping(context.Context) error
}
