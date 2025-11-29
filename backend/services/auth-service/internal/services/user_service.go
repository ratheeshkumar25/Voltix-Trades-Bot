package services

import (
	"errors"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
	repo "ratheeshkumar25/github.com/trading_bot/auth-service/internal/repo"
	"ratheeshkumar25/github.com/trading_bot/auth-service/pkg/hashpasscode"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExists        = errors.New("email already exists")
)

// UserService handles user-related operations
type UserService struct {
	DB   *gorm.DB
	Repo repo.AuthRepositoryInterface
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser creates a new user with trial subscription
func (s *UserService) CreateUser(email, passwords string, googleID *string) (*models.User, error) {
	// Check if email exists
	existingUser, _ := s.Repo.GetUserByEmail(email)
	if existingUser != nil {
		return nil, ErrEmailExists
	}

	// Hash password if provided
	var passwordHash string
	if passwords != "" {
		hash, err := hashpasscode.HashPassword(passwords)
		if err != nil {
			return nil, err
		}
		passwordHash = hash
	}

	// Create user
	user := &models.User{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         "user",
		IsActive:     true,
	}

	if googleID != nil {
		user.GoogleID = *googleID
	}

	// ensure the user has a UUID id before creating; ignore any numeric id returned by the repo
	user.ID = uuid.New()
	_, err := s.Repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Create 7-day trial subscription
	if err := s.Repo.CreateTrialSubscription(user.ID); err != nil {
		return nil, err
	}

	return user, nil
}
