package dbpostgres

import (
	"example-clean-go-application/internal/entities"
)

// GetUserByID returns a user by its ID using GORM
func (ur *Repository) GetUserByID(id string) (*entities.User, error) {
	user := &entities.User{}
	err := ur.DB.First(user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser creates a new user using GORM
func (ur *Repository) CreateUser(user *entities.User) error {
	err := ur.DB.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}