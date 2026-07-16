package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Handler struct {
	Logger *slog.Logger
	DB     *gorm.DB
}

func New(logger *slog.Logger, db *gorm.DB) *Handler {
	return &Handler{
		Logger: logger,
		DB:     db,
	}
}

func (h *Handler) Health(c fiber.Ctx) error {

	h.Logger.Info("Health check initiated")
	postgresDB, err := h.DB.DB()
	if err != nil {
		h.Logger.Error("Failed to get database connection", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to get database connection",
		})
	}

	if err := postgresDB.Ping(); err != nil {
		h.Logger.Error("Database ping failed", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":          "error",
			"message":         "Database ping failed",
			"database_status": "down",
		})
	}
	return c.JSON(fiber.Map{
		"status":          "ok",
		"database_status": "up",
		"service":         "barkbnb-api",
		"version":         "0.1.0",
	})
}
