package user

import (
	"example-clean-go-application/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Routes(r fiber.Router, uc *usecases.UserUseCase) {
	r.Get("/users/:id", func(c *fiber.Ctx) error {
		return GetUserByID(c, uc)
	})

	r.Post("/users", func(c *fiber.Ctx) error {
		return CreateUser(c, uc)
	})
}
