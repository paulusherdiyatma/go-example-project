// package services represents a service layer of the application
package services

import "example-layered-app/internal/repositories"

func (s *Service) GetProductByID(input GetProductByIDInput) (GetProductByIDOutput, error) {
	getProductByIDInput := repositories.GetProductByIDInput{
		ID: input.ID,
	}

	output, err := s.Repository.GetProductByID(getProductByIDInput)
	if err != nil {
		return GetProductByIDOutput{}, err
	}

	return GetProductByIDOutput{
		Product: (*Product)(output.Product),
	}, nil
}

func (s *Service) CreateProduct(input CreateProductInput) (CreateProductOutput, error) {
	createProductInput := repositories.CreateProductInput{
		Name:  input.Name,
		Price: input.Price,
	}

	_, err := s.Repository.CreateProduct(createProductInput)
	if err != nil {
		return CreateProductOutput{}, err
	}

	return CreateProductOutput{}, nil
}
