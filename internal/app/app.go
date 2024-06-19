package app

import (
	"github.com/wlcmtunknwndth/grpc_test/internal/app/grpcApp"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {

	GrpcApp := grpcApp.New(log, grpcPort)

	return &App{GRPCSrv: GrpcApp}
}
