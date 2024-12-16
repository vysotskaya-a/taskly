package app

import (
	"api-gateway/internal/closer"
	"api-gateway/internal/config"
	"api-gateway/internal/server/auth"
	"api-gateway/internal/server/project"
	"api-gateway/internal/server/task"
	"api-gateway/internal/server/user"
	authpb "api-gateway/pkg/api/auth_v1"
	projectpb "api-gateway/pkg/api/project_v1"
	taskpb "api-gateway/pkg/api/task_v1"
	userpb "api-gateway/pkg/api/user_v1"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	grpcConfig   config.GRPCConfig
	httpConfig   config.HTTPConfig
	loggerConfig config.LoggerConfig
	jwtConfig    config.JWTConfig

	userClientConn *grpc.ClientConn
	authAPIClient  authpb.AuthV1Client
	userAPIClient  userpb.UserV1Client

	projectClientConn *grpc.ClientConn
	projectAPIClient  projectpb.ProjectServiceClient
	taskAPIClient     taskpb.TaskServiceClient

	userHandler    *user.Handler
	authHandler    *auth.Handler
	projectHandler *project.Handler
	taskHandler    *task.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get grpc config: %w", err))
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get http config: %w", err))
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get logger config: %w", err))
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := config.NewJWTConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get jwt config: %w", err))
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}

func (s *serviceProvider) UserClientConn() *grpc.ClientConn {
	if s.userClientConn == nil {
		conn, err := grpc.NewClient(s.GRPCConfig().UserServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("failed to connect to gRPC user server: %w", err))
		}
		closer.Add(conn.Close)
		s.userClientConn = conn
	}
	return s.userClientConn
}

func (s *serviceProvider) AuthAPIClient() authpb.AuthV1Client {
	if s.authAPIClient == nil {
		s.authAPIClient = authpb.NewAuthV1Client(s.UserClientConn())
	}
	return s.authAPIClient
}

func (s *serviceProvider) UserAPIClient() userpb.UserV1Client {
	if s.userAPIClient == nil {
		s.userAPIClient = userpb.NewUserV1Client(s.UserClientConn())
	}
	return s.userAPIClient
}

func (s *serviceProvider) ProjectClientConn() *grpc.ClientConn {
	if s.projectClientConn == nil {
		conn, err := grpc.NewClient(s.GRPCConfig().ProjectServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("failed to connect to gRPC project server: %w", err))
		}
		closer.Add(conn.Close)
		s.projectClientConn = conn
	}
	return s.projectClientConn
}

func (s *serviceProvider) ProjectAPIClient() projectpb.ProjectServiceClient {
	if s.projectAPIClient == nil {
		s.projectAPIClient = projectpb.NewProjectServiceClient(s.ProjectClientConn())
	}
	return s.projectAPIClient
}

func (s *serviceProvider) TaskAPIClient() taskpb.TaskServiceClient {
	if s.taskAPIClient == nil {
		s.taskAPIClient = taskpb.NewTaskServiceClient(s.ProjectClientConn())
	}
	return s.taskAPIClient
}

func (s *serviceProvider) UserHandler() *user.Handler {
	if s.userHandler == nil {
		s.userHandler = user.NewHandler(s.UserAPIClient())
	}
	return s.userHandler
}

func (s *serviceProvider) AuthHandler() *auth.Handler {
	if s.authHandler == nil {
		s.authHandler = auth.NewHandler(s.AuthAPIClient(), s.JWTConfig())
	}
	return s.authHandler
}

func (s *serviceProvider) ProjectHandler() *project.Handler {
	if s.projectHandler == nil {
		s.projectHandler = project.NewHandler(s.ProjectAPIClient())
	}
	return s.projectHandler
}

func (s *serviceProvider) TaskHandler() *task.Handler {
	if s.taskHandler == nil {
		s.taskHandler = task.NewHandler(s.TaskAPIClient())
	}
	return s.taskHandler
}
