package repo

import (
	"github.com/google/uuid"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
	"gorm.io/gorm"
)

type AuthRepository struct {
	//Db connection and other fields as needed
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

// Interface for AuthRepository
type AuthRepositoryInterface interface {
	CreateUser(user *models.User) (uuid.UUID, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetUserByGoogleID(googleID string) (*models.User, error)
	AuthenticateUser(email, password string) (*models.User, error)
	DeleteUser(id uuid.UUID) error
	CreateTrialSubscription(userID uuid.UUID) error
	GetUserSubscription(userID uuid.UUID) (*models.Subscription, error)
	UpdateUserStatus(user *models.User) error
	CreateSession(session *models.Session) error
	AddExchangeCredential(cred *models.ExchangeCredential) error
	GetExchangeCredentials(userID uuid.UUID) ([]models.ExchangeCredential, error)
	UpdateExchangeCredential(cred *models.ExchangeCredential) error
	DeleteExchangeCredential(id uuid.UUID) error
}
