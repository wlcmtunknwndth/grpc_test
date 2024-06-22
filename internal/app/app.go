package app

import (
	"github.com/wlcmtunknwndth/grpc_test/internal/app/grpcApp"
	"github.com/wlcmtunknwndth/grpc_test/internal/services/auth"
	"github.com/wlcmtunknwndth/grpc_test/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	GrpcApp := grpcApp.New(log, grpcPort, authService)

	return &App{GRPCSrv: GrpcApp}
}
