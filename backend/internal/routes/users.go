package routes

import (
	"github.com/devenkulkarni/barkbnb/backend/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func registerUserRoutes(router fiber.Router, h *handlers.Handler) {
	router.Post("/users", h.CreateUser)
}
