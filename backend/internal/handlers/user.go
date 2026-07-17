package handlers

import (
	"github.com/devenkulkarni/barkbnb/backend/internal/dto"
	"github.com/devenkulkarni/barkbnb/backend/internal/models"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) CreateUser(c fiber.Ctx) error {
	req := dto.CreateUserRequest{}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := h.DB.Create(&user).Error; err != nil {

		h.Logger.Error("Failed to create user", "error", err)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
