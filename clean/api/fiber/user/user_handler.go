package user

import (
	"example-clean-go-application/internal/entities"
	"example-clean-go-application/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

func GetUserByID(c *fiber.Ctx, uc *usecases.UserUseCase) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User ID is required",
		})
	}

	user, err := uc.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx, uc *usecases.UserUseCase) error {
	requestBody := new(CreateUserRequest)
	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user := &entities.User{
		Username: requestBody.Username,
		Email:    requestBody.Email,
	}

	if err := uc.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}