package auth

import (
	"api-gateway/internal/errorz"
	"api-gateway/internal/models/request"
	"api-gateway/internal/models/response"
	"api-gateway/internal/server/helper"
	authpb "api-gateway/pkg/api/auth_v1"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetAccessToken(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var getAccessTokenRequest request.GetAccessToken
	if err := json.NewDecoder(r.Body).Decode(&getAccessTokenRequest); err != nil {
		return errorz.APIError{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("failed to decode GetAccessToken request body: %w", err),
			Msg:    "error decoding request body",
		}
	}

	getAccessTokenResp, err := h.authAPIClient.GetAccessToken(ctx, &authpb.GetAccessTokenRequest{
		RefreshToken: getAccessTokenRequest.RefreshToken,
	})
	if err != nil {
		return errorz.APIError{
			Status: http.StatusUnauthorized,
			Err:    fmt.Errorf("failed to get access token: %w", err),
			Msg:    "failed to get access token",
		}
	}

	resp := response.GetAccessToken{
		AccessToken: getAccessTokenResp.GetAccessToken(),
	}

	// Возвращаем ответ с токеном и статусом 200 (OK)
	return helper.WriteJSON(w, http.StatusOK, resp)
}
