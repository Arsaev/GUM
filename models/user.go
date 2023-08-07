package models

import "fmt"

// User represents a user
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// String returns a string representation of the User
func (u *User) String() string {
	return fmt.Sprintf("User<%d %s %s %s>", u.ID, u.FirstName, u.LastName, u.Email)
}
