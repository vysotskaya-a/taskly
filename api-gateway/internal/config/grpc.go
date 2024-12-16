package config

import (
	"errors"
	"net"
	"os"
)

var (
	errGRPCUserServerHostNotFound    = errors.New("grpc user server host not found")
	errGRPCUserServerPortNotFound    = errors.New("grpc user server port not found")
	errGRPCProjectServerHostNotFound = errors.New("grpc project server host not found")
	errGRPCProjectServerPortNotFound = errors.New("grpc project server port not found")
)

const (
	grpcUserServerHostEnvName    = "GRPC_USER_SERVER_HOST"
	grpcUserServerPortEnvName    = "GRPC_USER_SERVER_PORT"
	grpcProjectServerHostEnvName = "GRPC_PROJECT_SERVER_HOST"
	grpcProjectServerPortEnvName = "GRPC_PROJECT_SERVER_PORT"
)

type GRPCConfig interface {
	UserServerAddress() string
	ProjectServerAddress() string
}

type grpcConfig struct {
	userServerHost    string
	userServerPort    string
	projectServerHost string
	projectServerPort string
}

func NewGRPCConfig() (GRPCConfig, error) {
	userServerHost := os.Getenv(grpcUserServerHostEnvName)
	if len(userServerHost) == 0 {
		return nil, errGRPCUserServerHostNotFound
	}

	userServerPort := os.Getenv(grpcUserServerPortEnvName)
	if len(userServerPort) == 0 {
		return nil, errGRPCUserServerPortNotFound
	}

	projectServerHost := os.Getenv(grpcProjectServerHostEnvName)
	if len(projectServerHost) == 0 {
		return nil, errGRPCProjectServerHostNotFound
	}

	projectServerPort := os.Getenv(grpcProjectServerPortEnvName)
	if len(projectServerPort) == 0 {
		return nil, errGRPCProjectServerPortNotFound
	}

	return &grpcConfig{
		userServerHost:    userServerHost,
		userServerPort:    userServerPort,
		projectServerHost: projectServerHost,
		projectServerPort: projectServerPort,
	}, nil
}

func (cfg *grpcConfig) UserServerAddress() string {
	return net.JoinHostPort(cfg.userServerHost, cfg.userServerPort)
}

func (cfg *grpcConfig) ProjectServerAddress() string {
	return net.JoinHostPort(cfg.projectServerHost, cfg.projectServerPort)
}
