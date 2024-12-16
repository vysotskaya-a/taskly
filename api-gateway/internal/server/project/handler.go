package project

import projectpb "api-gateway/pkg/api/project_v1"

type Handler struct {
	projectAPIClient projectpb.ProjectServiceClient
}

func NewHandler(projectAPIClient projectpb.ProjectServiceClient) *Handler {
	return &Handler{projectAPIClient: projectAPIClient}
}
