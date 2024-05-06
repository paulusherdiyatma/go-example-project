
// dbpostgres package is an adapter to connect to the PostgreSQL database.
package dbpostgres


import "gorm.io/gorm"


// Repository is a struct that represents a database repository.
type Repository struct {
	DB *gorm.DB
}