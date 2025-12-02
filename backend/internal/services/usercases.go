package services

import (
	"github.com/google/uuid"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
	hashpasscode "github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/password"
)

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

// AuthenticateUser authenticates a user by email and password
func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !hashpasscode.CheckPassword(password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// CreateSession creates a new session
func (s *UserService) CreateSession(session *models.Session) error {
	return s.Repo.CreateSession(session)
}

// UpdateUser updates user details
func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	if err := s.Repo.UpdateUserStatus(user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(user *models.User) (*models.User, error) {
	if err := s.Repo.DeleteUser(user.ID); err != nil {
		return nil, err
	}
	return user, nil
}
