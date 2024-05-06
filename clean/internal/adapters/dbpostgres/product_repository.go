package dbpostgres

import (
	"example-clean-go-application/internal/entities"
)

func (pr *Repository) GetProductByID(id string) (*entities.Product, error) {
	product := &entities.Product{}
	err := pr.DB.First(product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *Repository) CreateProduct(product *entities.Product) error {
	err := pr.DB.Create(product).Error
	if err != nil {
		return err
	}

	return nil
}