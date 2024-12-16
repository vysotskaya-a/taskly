package project

import (
	"api-gateway/internal/errorz"
	"api-gateway/internal/models/response"
	"api-gateway/internal/server/helper"
	projectpb "api-gateway/pkg/api/project_v1"
	"fmt"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (h *Handler) GetProject(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	// Получаем id пользователя из url параметров
	projectID := chi.URLParam(r, "project_id")
	if len(projectID) == 0 {
		return errorz.APIError{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("failed to get project id"),
			Msg:    "failed to get project id",
		}
	}

	getProjectResp, err := h.projectAPIClient.GetProject(ctx, &projectpb.GetProjectRequest{
		Id: projectID,
	})
	if err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.Unauthenticated:
			return errorz.APIError{
				Status: http.StatusUnauthorized,
				Err:    err,
				Msg:    "authentication required",
			}
		case codes.PermissionDenied:
			return errorz.APIError{
				Status: http.StatusForbidden,
				Err:    err,
				Msg:    "permission denied",
			}
		case codes.NotFound:
			return errorz.APIError{
				Status: http.StatusNotFound,
				Err:    err,
				Msg:    "project not found",
			}
		case codes.Internal:
			return errorz.APIError{
				Status: http.StatusInternalServerError,
				Err:    err,
				Msg:    "failed to get project",
			}
		}
	}

	resp := response.GetProject{
		ID:                           getProjectResp.GetProject().GetId(),
		Title:                        getProjectResp.GetProject().GetTitle(),
		Description:                  getProjectResp.GetProject().GetDescription(),
		Users:                        getProjectResp.GetProject().GetUsers(),
		AdminID:                      getProjectResp.GetProject().GetAdminId(),
		NotificationSubscribersTGIds: getProjectResp.GetProject().GetNotificationSubscribersTgIds(),
		CreatedAt:                    getProjectResp.GetProject().GetCreatedAt().AsTime(),
	}

	return helper.WriteJSON(w, http.StatusOK, resp)
}
