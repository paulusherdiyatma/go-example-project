// package repositories represents a repository layer of the application
package repositories

import "gorm.io/gorm"

// NewRepository creates a new Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
