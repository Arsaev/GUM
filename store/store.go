package store

import "github.com/Arsaev/GUM/store/password"

type Client struct {
	// User     UserRepository
	Password password.Repository
}

func New() *Client {

}
