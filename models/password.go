package models

// Password represents a single password in the system for a user.
// User can have multiple passwords, but only one active password.
type Password struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	PasswordHash string `json:"passwordHash"`
	CreatedAt    int64  `json:"createdAt"`
}
