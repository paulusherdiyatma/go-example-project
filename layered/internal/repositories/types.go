package repositories

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

// GetUserByIDInput represents an input for GetUserByID function
type GetUserByIDInput struct {
	ID string
}

// GetUserByIDOutput represents an output for GetUserByID function
type GetUserByIDOutput struct {
	User *User
}

// GetProductByIDInput represents an input for GetProductByID function
type GetProductByIDInput struct {
	ID string
}

// GetProductByIDOutput represents an output for GetProductByID function
type GetProductByIDOutput struct {
	Product *Product
}

// CreateUserInput represents an input for CreateUser function
type CreateUserInput struct {
	ID       string
	Username string
	Email    string
}

// CreateUserOutput represents an output for CreateUser function
type CreateUserOutput struct {
	User *User
}

// CreateProductInput represents an input for CreateProduct function
type CreateProductInput struct {
	Name  string
	Price float64
}

// CreateProductOutput represents an output for CreateProduct function
type CreateProductOutput struct {
	Product *Product
}
