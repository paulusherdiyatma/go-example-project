// package repositories represents a repository layer of the application
package repositories

// GetProductByID queries a product by its ID
func (r *Repository) GetProductByID(input GetProductByIDInput) (GetProductByIDOutput, error) {
	r.DB.Where("id = ?", input.ID).First(&Product{})
	return GetProductByIDOutput{}, nil
}

// CreateProduct creates a new product
func (r *Repository) CreateProduct(input CreateProductInput) (CreateProductOutput, error) {
	r.DB.Create(&Product{
		Name:  input.Name,
		Price: input.Price,
	})
	return CreateProductOutput{}, nil
}
