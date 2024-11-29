package config

import (
	"errors"
	"os"
)

const (
	pgConnEnvName = "PG_CONN"
)

type PGConfig interface {
	PGConn() string
}

type pgConfig struct {
	pgConn string
}

func NewPGConfig() (PGConfig, error) {
	pgConn := os.Getenv(pgConnEnvName)
	if len(pgConn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		pgConn: pgConn,
	}, nil
}

func (cfg *pgConfig) PGConn() string {
	return cfg.pgConn
}
