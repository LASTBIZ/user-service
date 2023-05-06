package user

import (
	"context"
	"lastbiz/user-service/protos/user"
)

type router struct {
	service Service
	user.UnimplementedUserServiceServer
}

func NewUserRouter(service Service) user.UserServiceServer {
	return router{
		service: service,
	}
}

func (r router) GetUser(ctx context.Context, request *user.UserGetRequest) (*user.User, error) {

}

func (r router) CreateUser(ctx context.Context, u *user.User) (*user.User, error) {

}

func (r router) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*interface{}, error) {

}

func (r router) UpdateUser(ctx context.Context, u *user.User) (*user.User, error) {

}

func convertUserToGRPCUser(user User) *user.User {

}

func convertGRPUserToUser(authUser *user.User) User {

}
