package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/user"
	"net/http"
	"strings"
)

func GetUserByEmail(ctx context.Context, request *user.UserByEmailRequest, storage user1.Storage) (*user.UserResponse, error) {
	userEmail := request.GetEmail()
	if strings.TrimSpace(userEmail) == "" {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	u, err := storage.GetUserByEmail(userEmail)
	if err != nil {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusOK,
		User:   u.ToGRPC(),
	}, nil
}
