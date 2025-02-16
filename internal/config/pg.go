package config

import (
	"errors"
	"os"
)

type PgConfig interface {
	DSN() string
}

type PGConfig struct {
	Dsn string
}

func NewPgConfig() (PgConfig, error) {
	dsn := os.Getenv("DSN")
	if len(dsn) == 0 {
		return nil, errors.New("don't parse dsn")
	}
	return &PGConfig{
		Dsn: dsn,
	}, nil
}

func (c *PGConfig) DSN() string {
	return c.Dsn
}
