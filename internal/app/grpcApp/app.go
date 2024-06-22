package grpcApp

import (
	"context"
	"fmt"
	authgrpc "github.com/wlcmtunknwndth/grpc_test/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

const loc = "internal.app.grpcApp"

type Auth interface {
	Login(ctx context.Context,
		email string, password string, appID int,
	) (token string, err error)
	RegisterNewUser(ctx context.Context,
		email string, password string) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type App struct {
	log         *slog.Logger
	gRPCServer  *grpc.Server
	authService Auth
	port        int
}

func New(log *slog.Logger, port int, auth Auth) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer, auth)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}

}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = loc + ".Run"

	log := a.log.With(slog.String(op, op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = loc + ".Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
