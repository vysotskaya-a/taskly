package app

import (
	"context"
	"fmt"
	"user-service/internal/closer"
	"user-service/internal/config"
	"user-service/internal/repository"
	userRepository "user-service/internal/repository/postgres"
	authServer "user-service/internal/server/auth"
	userServer "user-service/internal/server/user"
	"user-service/internal/service"
	authService "user-service/internal/service/auth"
	userService "user-service/internal/service/user"

	"github.com/jmoiron/sqlx"

	db "user-service/pkg/pg"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	jwtConfig  config.JWTConfig

	db             *sqlx.DB
	userRepository repository.UserRepository

	userService service.UserService
	authService service.AuthService

	userServer *userServer.Server
	authServer *authServer.Server
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get pg config: %s", err.Error())) // лишняя паника и для типа err нужно использовать при форматировании флаг %w
			// fmt.Errorf("puk: %w", err)
			// в коде ниже то же самое 
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get grpc config: %s", err.Error())) // выше читать
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

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserServer(ctx context.Context) *userServer.Server {
	if s.userServer == nil {
		s.userServer = userServer.NewServer(s.UserService(ctx))
	}

	return s.userServer
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.UserRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AuthServer(ctx context.Context) *authServer.Server {
	if s.authServer == nil {
		s.authServer = authServer.NewServer(s.AuthService(ctx))
	}

	return s.authServer
}

// не совсем понял зачем нужны эти функции поясните плиз 