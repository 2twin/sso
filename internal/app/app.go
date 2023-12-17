package app

import (
	grpcapp "sso/internal/app/grpc"
	"time"

	"golang.org/x/exp/slog"
)

type App struct {
	GPRCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GPRCSrv: grpcApp,
	}
}
