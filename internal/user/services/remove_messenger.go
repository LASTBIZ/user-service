package services

import (
	"context"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/user"
	"net/http"
)

func RemoveMessenger(ctx context.Context, request *user.RemoveMessengerRequest, storage user1.Storage) (*user.RemoveMessengerResponse, error) {
	if request.GetMessengerId() == 0 {
		return &user.RemoveMessengerResponse{
			Status: http.StatusConflict,
			Error:  "id not found",
		}, nil
	}

	err := storage.RemoveMessenger(request.GetMessengerId(), request.GetUserId())

	if err != nil {
		return &user.RemoveMessengerResponse{
			Status: http.StatusConflict,
			Error:  "error remove messenger",
		}, nil
	}

	return &user.RemoveMessengerResponse{
		Status: http.StatusOK,
	}, nil
}
