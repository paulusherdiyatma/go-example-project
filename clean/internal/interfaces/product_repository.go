// package interfaces represents the interfaces of the application
package interfaces

import "example-clean-go-application/internal/entities"

// ProductRepository represents the interface that a product repository should implement
type ProductRepository interface {
	GetProductByID(id string) (*entities.Product, error)
	CreateProduct(product *entities.Product) error
}