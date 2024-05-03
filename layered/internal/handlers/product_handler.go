// package handlers represents a handler layer of the application
package handlers

import (
	"example-layered-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")
	if productID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid product id",
		})
	}

	output, err := h.Service.GetProductByID(services.GetProductByIDInput{
		ID: productID,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user": output.Product,
	})
}

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var input services.CreateProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	output, err := h.Service.CreateProduct(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user": output.Product,
	})
}
