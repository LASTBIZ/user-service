package user

import (
	"context"
	"lastbiz/user-service/pkg/logging"
	"lastbiz/user-service/pkg/user"
	"net/http"
	"strings"
	"time"
)

type Service struct {
	storage Storage
	user.UnimplementedUserServiceServer
}

func NewUserService(storage Storage) user.UserServiceServer {
	return Service{
		storage: storage,
	}
}

func (s Service) CreateUser(ctx context.Context, u *user.User) (*user.UserResponse, error) {
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

	if result, _ := s.storage.GetUserByEmail(u.Email); result != nil {
		return &user.UserResponse{
			Status: http.StatusConflict,
			Error:  "User already exists",
		}, nil
	}

	createUser := &User{
		Email:     u.Email,
		Lastname:  u.LastName,
		Firstname: u.FirstName,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.storage.CreateUser(*createUser)
	if err != nil {
		logger.Error(err)
		return &user.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error create user",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusCreated,
		User:   createUser.toGRPC(),
	}, nil
}

func (s Service) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.UserResponse, error) {
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

	err := s.storage.DeleteUser(deleteId)

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

func (s Service) UpdateUser(ctx context.Context, u *user.User) (*user.UserResponse, error) {
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

	if result, _ := s.storage.GetUserByEmail(u.Email); result == nil {
		return &user.UserResponse{
			Status: http.StatusConflict,
			Error:  "User not exists",
		}, nil
	}

	updateUser := &User{
		Blocked:   u.Blocked,
		Role:      u.Role,
		Firstname: u.FirstName,
		Lastname:  u.LastName,
		IsVerify:  u.IsVerify,
		Phone:     u.Phone,
	}

	updateUser, err := s.storage.UpdateUser(*updateUser)
	if err != nil {
		logger.Error(err)
		return &user.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error update user",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusOK,
		User:   updateUser.toGRPC(),
	}, nil
}

func (s Service) GetUser(ctx context.Context, request *user.UserGetRequest) (*user.UserResponse, error) {
	userId := request.GetUserId()
	if userId == 0 {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	u, err := s.storage.GetUser(userId)
	if err != nil {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusOK,
		User:   u.toGRPC(),
	}, nil
}

func (s Service) GetUserByEmail(ctx context.Context, request *user.UserByEmailRequest) (*user.UserResponse, error) {
	userEmail := request.GetEmail()
	if strings.TrimSpace(userEmail) == "" {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	u, err := s.storage.GetUserByEmail(userEmail)
	if err != nil {
		return &user.UserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusOK,
		User:   u.toGRPC(),
	}, nil
}
