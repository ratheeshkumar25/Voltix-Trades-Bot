package services

import (
	"errors"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
	repo "ratheeshkumar25/github.com/trading_bot/auth-service/internal/repo"
	"ratheeshkumar25/github.com/trading_bot/auth-service/pkg/hashpasscode"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExists        = errors.New("email already exists")
)

// UserServiceInterface defines the methods for user service
type UserServiceInterface interface {
	CreateUser(email, passwords string, googleID *string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetUserByGoogleID(googleID string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (*models.User, error)
	GetUserSubscription(userID uuid.UUID) (*models.Subscription, error)
	AuthenticateUser(email, password string) (*models.User, error)
	CreateSession(session *models.Session) error
	AddExchangeCredential(userID uuid.UUID, exchangeType models.ExchangeType, apiKey, apiSecret string) (*models.ExchangeCredential, error)
	GetExchangeCredentials(userID uuid.UUID) ([]models.ExchangeCredential, error)
	SwitchExchangeAccount(userID uuid.UUID, credentialID uuid.UUID) error
}

// UserService handles user-related operations
type UserService struct {
	Repo repo.AuthRepositoryInterface
}

// NewUserService creates a new user service
func NewUserService(repo repo.AuthRepositoryInterface) UserServiceInterface {
	return &UserService{
		Repo: repo,
	}
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

// AddExchangeCredential adds a new exchange credential
func (s *UserService) AddExchangeCredential(userID uuid.UUID, exchangeType models.ExchangeType, apiKey, apiSecret string) (*models.ExchangeCredential, error) {
	// TODO: Encrypt API Key and Secret
	cred := &models.ExchangeCredential{
		UserID:             userID,
		ExchangeType:       exchangeType,
		APIKeyEncrypted:    apiKey,    // Placeholder for encryption
		APISecretEncrypted: apiSecret, // Placeholder for encryption
		IsActive:           true,      // Default to active
	}
	if err := s.Repo.AddExchangeCredential(cred); err != nil {
		return nil, err
	}
	return cred, nil
}

// GetExchangeCredentials gets all exchange credentials for a user
func (s *UserService) GetExchangeCredentials(userID uuid.UUID) ([]models.ExchangeCredential, error) {
	return s.Repo.GetExchangeCredentials(userID)
}

// SwitchExchangeAccount switches the active exchange account
func (s *UserService) SwitchExchangeAccount(userID uuid.UUID, credentialID uuid.UUID) error {
	creds, err := s.Repo.GetExchangeCredentials(userID)
	if err != nil {
		return err
	}

	found := false
	for _, cred := range creds {
		if cred.ID == credentialID {
			cred.IsActive = true
			found = true
		} else {
			cred.IsActive = false
		}
		// Update all credentials to ensure only one is active (if that's the logic)
		// Or maybe we only update the ones that changed.
		// For simplicity, update all.
		if err := s.Repo.UpdateExchangeCredential(&cred); err != nil {
			return err
		}
	}

	if !found {
		return errors.New("credential not found")
	}

	return nil
}

// GetUserByID gets a user by ID
func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

// GetUserByGoogleID gets a user by Google ID
func (s *UserService) GetUserByGoogleID(googleID string) (*models.User, error) {
	return s.Repo.GetUserByGoogleID(googleID)
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	if err := s.Repo.UpdateUserStatus(user); err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user (placeholder implementation)
func (s *UserService) DeleteUser(user *models.User) (*models.User, error) {
	// TODO: Implement soft delete in Repo
	// For now just return nil
	return nil, nil
}

// GetUserSubscription gets a user's subscription
func (s *UserService) GetUserSubscription(userID uuid.UUID) (*models.Subscription, error) {
	return s.Repo.GetUserSubscription(userID)
}

// GetUserByEmail gets a user by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.Repo.GetUserByEmail(email)
}
