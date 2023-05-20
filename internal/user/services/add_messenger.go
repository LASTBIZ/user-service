package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/user"
	"net/http"
)

func AddMessenger(ctx context.Context, request *user.AddMessengerRequest, storage user1.Storage) (*user.AddMessengerResponse, error) {
	if request.GetMessenger() == nil {
		return &user.AddMessengerResponse{
			Status: http.StatusConflict,
			Error:  "messenger is null",
		}, nil
	}
	if request.GetUserId() == 0 {
		return &user.AddMessengerResponse{
			Status: http.StatusConflict,
			Error:  "user_id not found",
		}, nil
	}

	//check user is exits
	u, err := storage.GetUser(request.GetUserId())
	if err != nil {
		return &user.AddMessengerResponse{
			Status: http.StatusConflict,
			Error:  "user not found",
		}, nil
	}

	mrpc := request.GetMessenger()

	m := user1.Messenger{
		Name:   mrpc.GetName(),
		Value:  mrpc.GetValue(),
		UserId: u.ID,
	}

	m1, err := storage.AddMessenger(m)

	mrpc = &user.Messenger{
		Id:    m1.ID,
		Name:  m1.Name,
		Value: m1.Value,
	}

	return &user.AddMessengerResponse{
		Messenger: mrpc,
	}, nil
}
