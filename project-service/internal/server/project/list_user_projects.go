package project

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "project-service/pkg/api/project_v1"
)

func (s *Server) ListUserProjects(ctx context.Context, req *emptypb.Empty) (*pb.ListUserProjectsResponse, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Authentication required. Please provide a valid token.")
	}

	projects, err := s.projectService.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed get user projects.")
	}

	projectsPB := make([]*pb.Project, len(projects))
	for i, project := range projects {
		projectsPB[i] = &pb.Project{
			Id:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			Users:       project.Users,
			AdminId:     project.AdminID,
			CreatedAt:   timestamppb.New(project.CreatedAt),
		}
	}

	return &pb.ListUserProjectsResponse{
		Project: projectsPB,
	}, nil
}
