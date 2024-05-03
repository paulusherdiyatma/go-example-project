// package handlers represents a handler layer of the application
package handlers

import (
	"example-layered-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	output, err := h.Service.GetUserByID(services.GetUserByIDInput{
		ID: userID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user": output.User,
	})
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var input services.CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	output, err := h.Service.CreateUser(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user": output.User,
	})
}
