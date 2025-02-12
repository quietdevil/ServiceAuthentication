package pg

import (
	"context"
	"serviceauth/internal/client/db"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Key string

const TxKey Key = "tx"

type dbPg struct {
	dbc *pgxpool.Pool
}

func NewDBPG(pool *pgxpool.Pool) db.DB {
	return &dbPg{dbc: pool}
}

func (pg *dbPg) ExecContext(ctx context.Context, q db.Query, a ...any) (pgconn.CommandTag, error) {

	tx, ok := ctx.Value("tx").(pgx.Tx)
	if ok {
		tx.Query(ctx, q.QueryString, a...)
	}
	return pg.dbc.Exec(ctx, q.QueryString, a...)
}

func (pg *dbPg) QueryContext(ctx context.Context, q db.Query, a ...any) (pgx.Rows, error) {
	tx, ok := ctx.Value("tx").(pgx.Tx)
	if ok {
		tx.Query(ctx, q.QueryString, a...)
	}
	return pg.dbc.Query(ctx, q.QueryString, a...)
}

func (pg *dbPg) QueryRowContext(ctx context.Context, q db.Query, a ...any) pgx.Row {
	tx, ok := ctx.Value("tx").(pgx.Tx)
	if ok {
		tx.Query(ctx, q.QueryString, a...)
	}
	return pg.dbc.QueryRow(ctx, q.QueryString, a...)
}

func (pg *dbPg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return pg.dbc.BeginTx(ctx, txOptions)
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func (pg *dbPg) Close() {
	pg.dbc.Close()
}

func (pg *dbPg) Ping(ctx context.Context) error {
	return pg.dbc.Ping(ctx)
}
