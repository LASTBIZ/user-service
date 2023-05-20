package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/logging"
	"lastbiz/user-service/pkg/user"
	"net/http"
	"strings"
	"time"
)

func CreateUser(ctx context.Context, u *user.User, storage user1.Storage) (*user.UserResponse, error) {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"user": &u,
		"type": "createUser",
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

	if _, err := storage.GetUserByEmail(u.Email); err == nil {
		return &user.UserResponse{
			Status: http.StatusConflict,
			Error:  "User already exists",
		}, nil
	}

	createUser := &user1.User{
		Email:     u.Email,
		Lastname:  u.LastName,
		Firstname: u.FirstName,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := storage.CreateUser(*createUser)
	if err != nil {
		logger.Error(err)
		return &user.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error create user",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusCreated,
		User:   createUser.ToGRPC(),
	}, nil
}
