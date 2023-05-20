package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/logging"
	"lastbiz/user-service/pkg/user"
	"net/http"
)

func DeleteUser(ctx context.Context, request *user.DeleteUserRequest, storage user1.Storage) (*user.UserResponse, error) {
	deleteId := request.GetUserId()
	logger := logging.WithFields(ctx, map[string]interface{}{
		"userID": deleteId,
		"type":   "deleteUser",
	})
	if deleteId == 0 {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	err := storage.DeleteUser(deleteId)

	if err != nil {
		logger.Error(err)
		return &user.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error delete user",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusOK,
	}, nil
}
