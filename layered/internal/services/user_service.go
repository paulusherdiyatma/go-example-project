// package services represents a service layer of the application
package services

import (
	"example-layered-app/internal/repositories"
	"example-layered-app/pkg/utils"
)

func (s *Service) GetUserByID(input GetUserByIDInput) (GetUserByIDOutput, error) {
	getUserByIDInput := repositories.GetUserByIDInput{
		ID: input.ID,
	}

	output, err := s.Repository.GetUserByID(getUserByIDInput)
	if err != nil {
		return GetUserByIDOutput{}, err
	}

	return GetUserByIDOutput{
		User: (*User)(output.User),
	}, nil
}

func (s *Service) CreateUser(input CreateUserInput) (CreateUserOutput, error) {
	createUserInput := repositories.CreateUserInput{
		ID:       utils.GenerateUUID(),
		Email:    input.Email,
		Username: input.Username,
	}

	_, err := s.Repository.CreateUser(createUserInput)
	if err != nil {
		return CreateUserOutput{}, err
	}

	return CreateUserOutput{}, nil
}
