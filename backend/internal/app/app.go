package app

import (
	"log/slog"

	"github.com/devenkulkarni/barkbnb/backend/internal/config"
	"github.com/devenkulkarni/barkbnb/backend/internal/logger"
	"github.com/devenkulkarni/barkbnb/backend/internal/middleware"
	"github.com/devenkulkarni/barkbnb/backend/internal/routes"
	"github.com/gofiber/fiber/v3"
)

type Application struct {
	Config *config.Config
	Logger *slog.Logger
	Server *fiber.App
}

func NewApplication(cfg *config.Config) *Application {
	appLogger := logger.Loggerfunc(cfg)
	fiberApp := fiber.New()
	fiberApp.Use(
		middleware.RequestLogger(appLogger),
	)

	routes.Register(fiberApp)

	return &Application{
		Config: cfg,
		Server: fiberApp,
		Logger: appLogger,
	}
}

func (a *Application) Run() error {
	a.Logger.Info(
		"Starting HTTP server",
		"port", a.Config.Server.Port,
	)
	return a.Server.Listen(":" + a.Config.Server.Port)
}
