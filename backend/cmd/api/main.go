package main

import (
	"log"

	"github.com/devenkulkarni/barkbnb/backend/internal/routes"

	"github.com/gofiber/fiber/v3"

	"github.com/devenkulkarni/barkbnb/backend/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	app := fiber.New()

	routes.Register(app)

	log.Printf(
		"🐶 %s %s starting on port %s",
		cfg.App.Name,
		cfg.App.Version,
		cfg.Server.Port,
	)

	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
