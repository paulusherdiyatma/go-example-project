package main

import (
	"fmt"
	"os"

	"example-layered-app/internal/handlers"
	"example-layered-app/internal/middleware"
	"example-layered-app/internal/repositories"
	"example-layered-app/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("this is example layered app")
	db := initDB()

	repository := repositories.NewRepository(db)

	serviceOptions := services.NewServiceOptions{
		Repository: repository,
	}

	service := services.NewService(serviceOptions)
	initAPI(service)
}

func initDB() *gorm.DB {
	db, err := repositories.ConnectToDB()
	if err != nil {
		panic(err)
	}

	return db
}

func initAPI(service *services.Service) {
	app := fiber.New()
	app.Use(middleware.Logger)
	handlers.RegisterRoutes(app, service)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // default port
	}

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
