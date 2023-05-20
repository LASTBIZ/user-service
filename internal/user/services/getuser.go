package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/user"
	"net/http"
)

func GetUser(ctx context.Context, request *user.UserGetRequest, storage user1.Storage) (*user.UserResponse, error) {
	userId := request.GetUserId()
	if userId == 0 {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	u, err := storage.GetUser(userId)
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
