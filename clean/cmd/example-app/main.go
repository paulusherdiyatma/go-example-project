package main

import (
	apifiber "example-clean-go-application/api/fiber"
    "example-clean-go-application/internal/adapters/dbpostgres"
	"example-clean-go-application/internal/databases"
	"example-clean-go-application/internal/usecases"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	initFiberWithPostgres()
}

func initFiberWithPostgres() {
	app := fiber.New()

	db, err := database.ConnectPostgresDB()
	if err != nil {
		panic("could not connect to the database")
	}

	repo := dbpostgres.Repository{
		DB: db,
	}

	userUseCase := usecases.UserUseCase{
		UserRepository: &repo,
	}

	productUseCase := usecases.ProductUseCase{
		ProductRepository: &repo,
	}

	apifiber.RegisterRoutes(app, apifiber.UseCase{
		UserUseCase:    &userUseCase,
		ProductUseCase: &productUseCase,
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // default port
	}

	err = app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
