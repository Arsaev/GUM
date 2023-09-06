package gum

import (
	"fmt"
	"log"
	"time"

	"github.com/Arsaev/GUM/models"
	passwordStore "github.com/Arsaev/GUM/store/password"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	store passwordStore.Repository
}

func newPasswordService() *PasswordService {
	return &PasswordService{}
}

// Create implements PasswordRepository.Create method using bcrypt to hash the password.
func (s *PasswordService) Create(req *NewPassword) (*models.Password, error) {
	// Validate the request
	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("Failed to create password, request is invalid: %w", err)
	}

	// generate ID for the password
	id := uuid.New().String()

	// hash the password
	hash := hashAndSalt([]byte(req.Password))

	// create the password in the database
	pwd := &models.Password{
		ID:           id,
		UserID:       req.UserID,
		PasswordHash: hash,
		CreatedAt:    time.Now().Unix(),
	}

	err := s.store.Create(pwd)
	if err != nil {
		return nil, fmt.Errorf("Failed to create password: %w", err)
	}

	return pwd, nil
}

// Get implements PasswordRepository.Get method by getting the password from the database
// for the given user ID or password ID.
func (s *PasswordService) Get(req *GetPasswordRequest) (*models.Password, error) {
	// Validate the request
	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("Failed to get password, request is invalid: %w", err)
	}

	if req.ID != "" {
		pwd, err := s.store.GetByID(req.ID)
		if err != nil {
			return nil, fmt.Errorf("Failed to get password: %w", err)
		}

		return pwd, nil
	}

	pwds, err := s.store.GetByUserID(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get password: %w", err)
	}

	if len(pwds) == 0 {
		return nil, fmt.Errorf("Failed to get password: Password not found for the given user ID")
	}

	return pwds[0], nil
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("hashing and salting password failed with err %v", err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
