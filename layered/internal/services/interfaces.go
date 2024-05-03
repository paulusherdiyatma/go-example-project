// package services represents a service layer of the application
package services

import "example-layered-app/internal/repositories"

// Service struct represents a service
type Service struct {
	Repository *repositories.Repository
}

// ServiceInterface defines a contract for service layer
type ServiceInterface interface {
	GetUserByID(input GetUserByIDInput) (GetUserByIDOutput, error)
	GetProductByID(input GetProductByIDInput) (GetProductByIDOutput, error)

	CreateUser(input CreateUserInput) (CreateUserOutput, error)
	CreateProduct(input CreateProductInput) (CreateProductOutput, error)
}
