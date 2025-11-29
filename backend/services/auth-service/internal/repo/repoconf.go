package repo

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"

	"github.com/google/uuid"
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

//Interface for AuthRepository

type AuthRepositoryInterface interface {
	CreateUser(user *models.User) (uint, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetUserByGoogleID(googleID string) (*models.User, error)
	AuthenticateUser(email, password string) (*models.User, error)
	CreateTrialSubscription(userID uuid.UUID) error
	GetUserSubscription(userID uuid.UUID) (*models.Subscription, error)
	UpdateUserStatus(user *models.User) error
}
