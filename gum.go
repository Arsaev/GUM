package gum

// Gum is the main struct that implements the Gum interface
type Gum struct {
	User     UserRepository     // User repository
	Password PasswordRepository // Password repository
}

// New returns a new Gum instance
func New() *Gum {
	g := &Gum{}
	return g
}

func (g *Gum) setup(param string) error {
	g.User = newUserService()
	return nil
}
