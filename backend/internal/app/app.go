package app

import (
	"github.com/gofiber/fiber/v3"

	"github.com/devenkulkarni/barkbnb/backend/internal/config"
	"github.com/devenkulkarni/barkbnb/backend/internal/routes"
)

type Application struct {
	Config *config.Config
	Server *fiber.App
}

func NewApplication(cfg *config.Config) *Application {
	fiberApp := fiber.New()
	routes.Register(fiberApp)

	return &Application{
		Config: cfg,
		Server: fiberApp,
	}
}

func (a *Application) Run() error {
	return a.Server.Listen(":" + a.Config.Server.Port)
}
