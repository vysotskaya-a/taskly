package task

import (
	"api-gateway/internal/errorz"
	"api-gateway/internal/models/request"
	"api-gateway/internal/models/response"
	"api-gateway/internal/server/helper"
	taskpb "api-gateway/pkg/api/task_v1"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) error {
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

	var createTaskRequest request.CreateTask
	if err := json.NewDecoder(r.Body).Decode(&createTaskRequest); err != nil {
		return errorz.APIError{
			Status: http.StatusBadRequest,
			Err:    err,
			Msg:    "failed to decode create task request",
		}
	}

	createTaskResp, err := h.taskAPIClient.CreateTask(ctx, &taskpb.CreateTaskRequest{
		Title:       createTaskRequest.Title,
		Description: createTaskRequest.Description,
		Status:      createTaskRequest.Status,
		ProjectId:   projectID,
		Executor:    createTaskRequest.Executor,
		Deadline:    timestamppb.New(createTaskRequest.Deadline),
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
				Msg:    "failed to create task",
			}
		}
	}

	resp := response.CreateTask{
		ID: createTaskResp.GetId(),
	}

	return helper.WriteJSON(w, http.StatusCreated, resp)
}
