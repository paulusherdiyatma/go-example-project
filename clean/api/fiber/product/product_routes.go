package product

import (
	"example-clean-go-application/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

// Routes defines the product routes
func Routes(r fiber.Router, uc *usecases.ProductUseCase) {
	r.Get("/products/:id", func(c *fiber.Ctx) error {
		return GetProductByID(c, uc)
	})
}