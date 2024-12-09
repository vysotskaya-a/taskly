package app

import (
	"context"
	"fmt"
	"project-service/internal/repository"
	"project-service/internal/service"

	"github.com/jmoiron/sqlx"

	projectRepository "project-service/internal/repository/postgres"

	projectservice "project-service/internal/service/project"

	projectServer "project-service/internal/server/project"
	db "project-service/pkg/pg"

	"project-service/internal/closer"
	"project-service/internal/config"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	db                *sqlx.DB
	projectRepository repository.ProjectRepository

	projectService service.ProjectService

	projectServer *projectServer.Server
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
		database := db.Init(s.PGConfig().PGConn())
		closer.Add(database.Close)
		s.db = database
	}

	return s.db
}

func (s *serviceProvider) ProjectRepository(ctx context.Context) repository.ProjectRepository {
	if s.projectRepository == nil {
		s.projectRepository = projectRepository.NewProjectRepository(s.DBClient(ctx))
	}

	return s.projectRepository
}

func (s *serviceProvider) ProjectService(ctx context.Context) service.ProjectService {
	if s.projectService == nil {
		s.projectService = projectservice.NewService(s.ProjectRepository(ctx))
	}

	return s.projectService
}

func (s *serviceProvider) ProjectServer(ctx context.Context) *projectServer.Server {
	if s.projectServer == nil {
		s.projectServer = projectServer.NewServer(s.ProjectService(ctx))
	}

	return s.projectServer
}
