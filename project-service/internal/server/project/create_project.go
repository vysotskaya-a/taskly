package project

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-service/internal/models"
	pb "project-service/pkg/api/project_v1"
)

func (s *Server) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Authentication required. Please provide a valid token.")
	}

	project := &models.Project{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Users:       []string{userID},
		AdminID:     userID,
	}

	projectID, err := s.projectService.Create(ctx, project)
	if err != nil {
		log.Error().Err(err).Msg("error while creating project")
		return nil, status.Error(codes.Internal, "Failed to create project.")
	}

	return &pb.CreateProjectResponse{
		Id: projectID,
	}, nil
}
