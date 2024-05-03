// package services represents a service layer of the application
package services

// User represents a user entity
type User struct {
	ID       string
	Username string
	Email    string
}

// Product represents a product entity
type Product struct {
	ID    string
	Name  string
	Price float64
}

type GetUserByIDInput struct {
	ID string
}

type GetUserByIDOutput struct {
	User *User
}

type GetProductByIDInput struct {
	ID string
}

type GetProductByIDOutput struct {
	Product *Product
}

type CreateUserInput struct {
	Username string
	Email    string
}

type CreateUserOutput struct {
	User *User
}

type CreateProductInput struct {
	Name  string
	Price float64
}

type CreateProductOutput struct {
	Product *Product
}
