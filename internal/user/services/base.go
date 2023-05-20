package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/user"
)

type Service struct {
	storage user1.Storage
	user.UnimplementedUserServiceServer
}

func NewUserService(storage user1.Storage) user.UserServiceServer {
	return Service{
		storage: storage,
	}
}

func (s Service) CreateUser(ctx context.Context, u *user.User) (*user.UserResponse, error) {
	return CreateUser(ctx, u, s.storage)
}

func (s Service) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.UserResponse, error) {
	return DeleteUser(ctx, request, s.storage)
}

func (s Service) UpdateUser(ctx context.Context, u *user.User) (*user.UserResponse, error) {
	return UpdateUser(ctx, u, s.storage)
}

func (s Service) GetUser(ctx context.Context, request *user.UserGetRequest) (*user.UserResponse, error) {
	return GetUser(ctx, request, s.storage)
}

func (s Service) GetUserByEmail(ctx context.Context, request *user.UserByEmailRequest) (*user.UserResponse, error) {
	return GetUserByEmail(ctx, request, s.storage)
}

func (s Service) AddMessenger(ctx context.Context, req *user.AddMessengerRequest) (*user.AddMessengerResponse, error) {
	return AddMessenger(ctx, req, s.storage)
}

func (s Service) RemoveMessenger(ctx context.Context, req *user.RemoveMessengerRequest) (*user.RemoveMessengerResponse, error) {
	return RemoveMessenger(ctx, req, s.storage)
}
