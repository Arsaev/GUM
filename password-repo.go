package gum

import (
	"fmt"

	"github.com/Arsaev/GUM/models"
)

// PasswordRepository represents a repository for managing passwords.
type PasswordRepository interface {
	Create(req *NewPassword) (*models.Password, error)
	Get(req *GetPasswordRequest) (*models.Password, error)
}

// NewPassword is request body for creating a new password.
type NewPassword struct {
	UserID   string `json:"userId"`   // ID of the user (required)
	Password string `json:"password"` // plain text password (required)
}

func (np *NewPassword) validate() error {
	if np.UserID == "" {
		return fmt.Errorf("UserID is required to create a new password")
	}

	if np.Password == "" {
		return fmt.Errorf("Password is required to create a new password")
	}
	return nil
}

// GetPasswordRequest is request body for getting a password.
// Either ID or UserID is required.
type GetPasswordRequest struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

func (gpr *GetPasswordRequest) validate() error {
	if gpr.ID == "" && gpr.UserID == "" {
		return fmt.Errorf("Either ID or UserID is required to get a password")
	}
	return nil
}
