package app

import (
	"log/slog"

	"github.com/devenkulkarni/barkbnb/backend/internal/config"
	"github.com/devenkulkarni/barkbnb/backend/internal/database"
	"github.com/devenkulkarni/barkbnb/backend/internal/handlers"
	"github.com/devenkulkarni/barkbnb/backend/internal/logger"
	"github.com/devenkulkarni/barkbnb/backend/internal/middleware"
	"github.com/devenkulkarni/barkbnb/backend/internal/routes"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Application struct {
	Config   *config.Config
	Logger   *slog.Logger
	Server   *fiber.App
	Database *gorm.DB
}

func NewApplication(cfg *config.Config) (*Application, error) {
	appLogger := logger.Loggerfunc(cfg)
	fiberApp := fiber.New()
	fiberApp.Use(
		middleware.RequestLogger(appLogger),
	)

	dbApp, err := database.NewDB(cfg)
	if err != nil {
		appLogger.Error("Failed to connect to the database", "error", err)

		return nil, err
	}

	appLogger.Info(
		"Connected to PostgreSQL",
		"host", cfg.Database.Host,
		"port", cfg.Database.Port,
		"database", cfg.Database.Name,
	)

	handler := handlers.New(appLogger, dbApp)
	routes.Register(fiberApp, handler)

	return &Application{
		Config:   cfg,
		Server:   fiberApp,
		Logger:   appLogger,
		Database: dbApp,
	}, nil
}

func (a *Application) Run() error {
	a.Logger.Info(
		"Starting HTTP server",
		"port", a.Config.Server.Port,
	)
	return a.Server.Listen(":" + a.Config.Server.Port)
}
