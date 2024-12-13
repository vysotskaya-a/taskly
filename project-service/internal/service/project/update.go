package project

import (
	"context"
	"project-service/internal/errorz"
	"project-service/internal/models"
	pb "project-service/pkg/api/project_v1"
)

func (s *Service) Update(ctx context.Context, req *pb.UpdateProjectRequest) error {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return errorz.ErrUserIDNotSet
	}

	project, err := s.projectRepository.GetByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if userID != project.AdminID {
		return errorz.ErrProjectAccessForbidden
	}

	s.update(project, req)

	return s.projectRepository.Update(ctx, project)
}

func (s *Service) update(project *models.Project, req *pb.UpdateProjectRequest) {
	if req.GetTitle() != "" {
		project.Title = req.GetTitle()
	}
	if req.GetDescription() != "" {
		project.Description = req.GetDescription()
	}
	if req.GetAdminId() != "" {
		project.AdminID = req.GetAdminId()
	}
}
