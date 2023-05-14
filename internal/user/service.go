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
	log     logging.Logger
	user.UnimplementedUserServiceServer
}

func NewUserService() user.UserServiceServer {
	return Service{}
}

func (s Service) CreateUser(ctx context.Context, u *user.User) (*user.UserResponse, error) {
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
		return &user.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error create user",
		}, nil
	}

	return &user.UserResponse{
		Status: http.StatusCreated,
		//User:   createUser,
	}, nil
}

func (s Service) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.UserResponse, error) {
	//return s.storage.DeleteUser(userId)
	return nil, nil
}

func (s Service) UpdateUser(ctx context.Context, u *user.User) (*user.UserResponse, error) {
	//return s.storage.UpdateUser(user)
	return nil, nil
}

func (s Service) GetUser(ctx context.Context, request *user.UserGetRequest) (*user.UserResponse, error) {
	return nil, nil
}

func (s Service) GetUserByEmail(ctx context.Context, request *user.UserByEmailRequest) (*user.UserResponse, error) {
	return nil, nil
}
