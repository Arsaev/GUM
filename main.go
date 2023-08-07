package gum

import (
	"github.com/Arsaev/GUM/logic"
	"github.com/Arsaev/GUM/models"
)

type Gum struct {
	User UserService
}

type UserService interface {
	Get(id string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
	List(query *logic.ListQuery) ([]*models.User, error)
}

// New returns a new Gum instance
func New() *Gum {

	g := &Gum{}

	userService := logic.NewUserService()
	g.User = userService

	return g
}
