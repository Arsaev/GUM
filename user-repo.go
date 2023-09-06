package gum

import "fmt"

// UserRepository represents a repository for managing users.
type UserRepository interface {
	Create(req *NewUser) (*User, error)
	Update(req *User) error
	Get(req *GetUserRequest) (*User, error)
}

// User represents a single user in the system.
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// NewUser is request body for creating a new user.
type NewUser struct {
	Email    string `json:"email"`    // Email address of the user (required)
	Password string `json:"password"` // Password of the user (required)

	FirstName string `json:"first_name"` // First name of the user (optional)
	LastName  string `json:"last_name"`  // Last name of the user (optional)
}

func (nu *NewUser) validate() error {
	return fmt.Errorf("not implemented")
}

// GetUserRequest is request body for getting a user.
// Either ID or Email is required.
type GetUserRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (gur *GetUserRequest) validate() error {
	return fmt.Errorf("not implemented")
}
