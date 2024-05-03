// package handlers represents a handler layer of the application
package handlers

import "example-layered-app/internal/services"

type Handler struct {
	Service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
