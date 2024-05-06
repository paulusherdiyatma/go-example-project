// package usecases represents the use cases of the application
package usecases

import (
	"example-clean-go-application/internal/entities"
	"example-clean-go-application/internal/interfaces"
)

// ProductUseCase contains functions related to products
type ProductUseCase struct {
	ProductRepository interfaces.ProductRepository
}

// GetProductByID returns a product by its ID
func (uc *ProductUseCase) GetProductByID(id string) (*entities.Product, error) {
	product, err := uc.ProductRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// CreateProduct creates a new product
func (uc *ProductUseCase) CreateProduct(product *entities.Product) error {
	err := uc.ProductRepository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}
