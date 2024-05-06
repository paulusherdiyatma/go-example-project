// package usecases represents the use cases of the application
package usecases

import (
	"example-clean-go-application/internal/entities"
	"example-clean-go-application/internal/interfaces"
)

// UserUseCase contains functions related to users
type UserUseCase struct {
	UserRepository interfaces.UserRepository
}

// GetUserByID returns a user by its ID
func (uc *UserUseCase) GetUserByID(id string) (*entities.User, error) {
	user, err := uc.UserRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(user *entities.User) error {
	err := uc.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}