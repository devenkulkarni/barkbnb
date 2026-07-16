package routes

import (
	"github.com/devenkulkarni/barkbnb/backend/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func Register(app *fiber.App, h *handlers.Handler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", h.Health)
	// registerHealthRoutes(v1, h)
}
