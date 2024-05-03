// package repositories represents a repository layer of the application
package repositories

import "gorm.io/gorm"

// Repository struct contains a pointer to a gorm.DB instance
type Repository struct {
	DB *gorm.DB
}

// RepositoryReader interface contains functions for querying
type RepositoryReader interface {
	GetUserByID(input GetUserByIDInput) (GetUserByIDOutput, error)
	GetProductByID(input GetProductByIDInput) (GetProductByIDOutput, error)
}

// RepositoryWriter interface contains functions for managing data
type RepositoryWriter interface {
	CreateUser(input CreateUserInput) (CreateUserOutput, error)
	CreateProduct(input CreateProductInput) (CreateProductOutput, error)
}
