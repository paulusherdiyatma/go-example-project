package apifiber

import (
	"example-clean-go-application/api/fiber/product"
	"example-clean-go-application/api/fiber/user"
	"example-clean-go-application/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type UseCase struct {
	UserUseCase    *usecases.UserUseCase
	ProductUseCase *usecases.ProductUseCase
}

func RegisterRoutes(app *fiber.App, useCase UseCase) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user.Routes(v1, useCase.UserUseCase)
	product.Routes(v1, useCase.ProductUseCase)
}