package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/logging"
	"lastbiz/user-service/pkg/user"
	"net/http"
	"strings"
)

func UpdateUser(ctx context.Context, u *user.User, storage user1.Storage) (*user.UserResponse, error) {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"user": &u,
		"type": "updateUser",
	})
	if strings.TrimSpace(u.Email) == "" {
		return &user.UserResponse{
			Status: http.StatusConflict,
			Error:  "email is empty",
		}, nil
	}

	if strings.TrimSpace(u.FirstName) == "" {
		return &user.UserResponse{
			Status: http.StatusConflict,
			Error:  "FirstName is empty",
		}, nil
	}

	if strings.TrimSpace(u.LastName) == "" {
		return &user.UserResponse{
			Status: http.StatusConflict,
			Error:  "LastName is empty",
		}, nil
	}

	if result, _ := storage.GetUserByEmail(u.Email); result == nil {
		return &user.UserResponse{
			Status: http.StatusConflict,
			Error:  "User not exists",
		}, nil
	}

	updateUser := &user1.User{
		ID:        u.Id,
		Blocked:   u.Blocked,
		Role:      u.Role,
		Firstname: u.FirstName,
		Lastname:  u.LastName,
		IsVerify:  u.IsVerify,
		Phone:     u.Phone,
	}

	updateUser, err := storage.UpdateUser(*updateUser)
	if err != nil {
		logger.Error(err)
		return &user.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error update user",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusOK,
		User:   updateUser.ToGRPC(),
	}, nil
}
