package pg

import (
	"context"
	"serviceauth/internal/client/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgClient struct {
	pgClienDB db.DB
}

func NewPgClient(ctx context.Context, dsn string) (db.Client, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &PgClient{pgClienDB: &dbPg{dbc: pool}}, nil
}

func (pg *PgClient) Close() error {
	pg.DB().Close()
	return nil

}

func (pg *PgClient) DB() db.DB {
	return pg.pgClienDB
}
