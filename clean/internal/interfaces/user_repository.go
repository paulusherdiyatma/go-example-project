// package interfaces represents the interfaces of the application
package interfaces

import "example-clean-go-application/internal/entities"

// UserRepository represents the interface that a user repository should implement
type UserRepository interface {
	GetUserByID(id string) (*entities.User, error)
	CreateUser(user *entities.User) error
}