// package services represents a service layer of the application
package services

import "example-layered-app/internal/repositories"

type NewServiceOptions struct {
	Repository *repositories.Repository
}

func NewService(opt NewServiceOptions) *Service {
	return &Service{
		Repository: opt.Repository,
	}
}
