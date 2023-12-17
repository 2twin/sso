package main

import (
	// "fmt"
	"os"
	"sso/internal/app"
	"sso/internal/config"

	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	
	log := setupLogger(cfg.Env)

	log.Info(
		"starting application", 
		slog.Any("cinfig", cfg),
	)

	application := app.New(log, cfg.GPRC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.GPRCSrv.MustRun()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout, 
				&slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		)		
	}

	return log
}
