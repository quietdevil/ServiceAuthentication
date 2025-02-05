package config

import (
	"errors"
	"os"
)

type PGConfig struct {
	DSN string
}

func NewPgConfig() (*PGConfig, error) {
	dsn := os.Getenv("DSN")
	if len(dsn) == 0 {
		return nil, errors.New("don't parse dsn")
	}
	return &PGConfig{
		DSN: dsn,
	}, nil
}

func (c *PGConfig) GetDSN() string {
	return c.DSN
}
