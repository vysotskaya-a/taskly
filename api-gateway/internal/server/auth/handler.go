package auth

import (
	"api-gateway/internal/config"
	authpb "api-gateway/pkg/api/auth_v1"
)

type Handler struct {
	authAPIClient authpb.AuthV1Client

	jwtConfig config.JWTConfig
}

func NewHandler(authAPIClient authpb.AuthV1Client, jwtConfig config.JWTConfig) *Handler {
	return &Handler{authAPIClient: authAPIClient, jwtConfig: jwtConfig}
}
