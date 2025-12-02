package services

import (
	"errors"

	"github.com/google/uuid"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/repo"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/http"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExists        = errors.New("email already exists")
)

type UserService struct {
	Repo repo.AuthRepositoryInterface
	App  *http.App
}

// GetUserByEmail implements UserServiceInterface.
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByGoogleID implements UserServiceInterface.
func (s *UserService) GetUserByGoogleID(googleID string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByID implements UserServiceInterface.
func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	panic("unimplemented")
}

// GetUserSubscription implements UserServiceInterface.
func (s *UserService) GetUserSubscription(userID uuid.UUID) (*models.Subscription, error) {
	panic("unimplemented")
}

func NewUserService(app *http.App, repo repo.AuthRepositoryInterface) *UserService {
	return &UserService{
		Repo: repo,
		App:  app,
	}
}

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
