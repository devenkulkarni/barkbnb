package routes

import (
	"github.com/devenkulkarni/barkbnb/backend/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func registerHealthRoutes(router fiber.Router) {
	router.Get("/health", handlers.Health)
}
