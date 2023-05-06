package user

import (
	"lastbiz/user-service/pkg/logging"
)

type Service struct {
	storage Storage
	log     logging.Logger
}

func NewUserService() Service {
	return Service{}
}

func (s Service) CreateUser(user User) (User, error) {
	return s.storage.CreateUser(user)
}

func (s Service) DeleteUser(userId uint32) error {
	return s.storage.DeleteUser(userId)
}

func (s Service) UpdateUser(user User) (User, error) {
	return s.storage.UpdateUser(user)
}

func (s Service) GetUser(userId uint32) (User, error) {
	return s.storage.GetUser(userId)
}

func (s Service) GetUserByEmail(email string) (User, error) {
	return s.storage.GetUserByEmail(email)
}
