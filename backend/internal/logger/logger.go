package logger

import (
	"log/slog"
	"os"

	"github.com/devenkulkarni/barkbnb/backend/internal/config"
)

func Loggerfunc(cfg *config.Config) *slog.Logger {

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	Loggerfunc := slog.New(handler)

	return Loggerfunc.With(
		"service", cfg.App.Name,
		"version", cfg.App.Version,
		"environment", cfg.App.Env,
	)

}
