package product

import (
	"example-clean-go-application/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

// GetProductByID returns a product by its ID
func GetProductByID(c *fiber.Ctx, uc *usecases.ProductUseCase) error {
	productID := c.Params("id")
	if productID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Product ID is required",
		})
	}

	product, err := uc.GetProductByID(productID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return c.JSON(product)
}