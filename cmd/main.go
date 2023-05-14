package main

import (
	"context"
	"google.golang.org/grpc"
	"lastbiz/user-service/internal/config"
	user1 "lastbiz/user-service/internal/user"
	"lastbiz/user-service/pkg/logging"
	"lastbiz/user-service/pkg/postgres"
	"lastbiz/user-service/pkg/user"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logging.Info(ctx, "config initializing")
	cfg := config.GetConfig()

	pgconfig := postgres.NewPGConfig(cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DB)

	pgClient := postgres.NewClient(ctx, 5, time.Second*5, pgconfig)
	userStorage := user1.NewUserStorage(*pgClient)
	userService := user1.NewUserService(userStorage)

	logging.Info(ctx, "run application")
	lis, err := net.Listen("tcp", cfg.GRPCPort)

	if err != nil {
		logging.GetLogger().Fatal(err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userService)

	if err := grpcServer.Serve(lis); err != nil {
		logging.GetLogger().Fatal(err)
	}
}
