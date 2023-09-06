package password

import "github.com/Arsaev/GUM/models"

// Repository represents a repository for managing passwords.
type Repository interface {
	Create(req *models.Password) error
	// GetByUserID returns a slice of passwords for the given user ID.
	// sorted by created_at field in descending order.
	GetByUserID(userID string) ([]*models.Password, error)
	// returns last 100 passwords.
	GetByID(id string) (*models.Password, error)
}
