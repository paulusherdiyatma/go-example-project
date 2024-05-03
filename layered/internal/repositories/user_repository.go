// package repositories represents a repository layer of the application
package repositories

// GetUserByID queries a user by its ID
func (r *Repository) GetUserByID(input GetUserByIDInput) (GetUserByIDOutput, error) {
	r.DB.Where("id = ?", input.ID).First(&User{})
	return GetUserByIDOutput{}, nil
}

// CreateUser creates a new user
func (r *Repository) CreateUser(input CreateUserInput) (CreateUserOutput, error) {
	r.DB.Create(&User{
		ID:       input.ID,
		Email:    input.Email,
		Username: input.Username,
	})
	return CreateUserOutput{}, nil
}
