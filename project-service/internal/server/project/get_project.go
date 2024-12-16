package project

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"project-service/internal/errorz"
	pb "project-service/pkg/api/project_v1"
)

func (s *Server) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	project, err := s.projectService.GetByID(ctx, req.GetId())
	switch {
	case errors.Is(err, errorz.ErrUserIDNotSet):
		return nil, status.Error(codes.Unauthenticated, "Authentication required. Please provide a valid token.")

	case errors.Is(err, errorz.ErrProjectNotFound):
		return nil, status.Error(codes.NotFound, "Project with this ID not found.")

	case errors.Is(err, errorz.ErrProjectAccessForbidden):
		return nil, status.Error(codes.PermissionDenied, "Access denied for this project.")

	case err != nil:
		log.Error().Err(err).Msg("error while getting project")
		return nil, status.Error(codes.Internal, "Failed to get project.")
	}

	return &pb.GetProjectResponse{
		Project: &pb.Project{
			Id:                           project.ID,
			Title:                        project.Title,
			Description:                  project.Description,
			Users:                        project.Users,
			AdminId:                      project.AdminID,
			NotificationSubscribersTgIds: project.NotificationSubscribersTGIDS,
			CreatedAt:                    timestamppb.New(project.CreatedAt),
		},
	}, nil
}
