package app

import (
	"context"
	"fmt"
	"user-service/internal/closer"
	"user-service/internal/config"

	"github.com/jmoiron/sqlx"

	db "user-service/pkg/pg"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	db *sqlx.DB
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get pg config: %s", err.Error()))
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get grpc config: %s", err.Error()))
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) *sqlx.DB {
	if s.db == nil {
		database := db.Init(s.pgConfig.PGConn())
		closer.Add(database.Close)
		s.db = database
	}

	return s.db
}
