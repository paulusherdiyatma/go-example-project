// package handlers represents a handler layer of the application
package handlers

import (
	"example-layered-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, service *services.Service) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	registerUserRoutes(v1, service)
	registerProductRoutes(v1, service)

	app.Group("/api/v1", func(ctx *fiber.Ctx) error {
		return nil
	})
}

func registerUserRoutes(r fiber.Router, service *services.Service) {
	userHandler := NewHandler(service)
	r.Get("/users/:id", userHandler.GetUserByID)
	r.Post("/users", userHandler.CreateUser)
}

func registerProductRoutes(r fiber.Router, service *services.Service) {
	productHandler := NewHandler(service)
	r.Get("/products/:id", productHandler.GetProductByID)
	r.Post("/products", productHandler.CreateProduct)
}
