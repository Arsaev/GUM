package gum

import (
	"fmt"
)

// UserService implements UserRepository interface.
type UserService struct {
	Password PasswordRepository
}

type userServiceConfig struct {
	Password PasswordRepository
}

// NewUserService returns a new instance of UserService.
func newUserService(cfg *userServiceConfig) *UserService {
	return &UserService{
		Password: cfg.Password,
	}
}

// Create creates a new user.
func (s *UserService) Create(req *NewUser) (*User, error) {
	// Validate the request
	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("Failed to create user, request is invalid: %w", err)
	}

	// generate ID for the user

	// create the user in the database

	// hash the password of the user and store it in the database in different table

	// return the user
	return nil, fmt.Errorf("not implemented")
}

// Update updates an existing user.
func (s *UserService) Update(req *User) error {
	// Validate the request

	// check if the user exists

	// create the user in the database

	// return the user
	return fmt.Errorf("not implemented")
}

func (s *UserService) Get(req *GetUserRequest) (*User, error) {
	// Validate the request
	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("Failed to get user, request is invalid: %w", err)
	}

	return nil, fmt.Errorf("not implemented")
}
