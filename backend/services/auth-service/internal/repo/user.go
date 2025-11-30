package repo

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
	"ratheeshkumar25/github.com/trading_bot/auth-service/pkg/hashpasscode"

	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Create user
func (r *AuthRepository) CreateUser(user *models.User) (uuid.UUID, error) {
	result := r.DB.Create(user)
	return user.ID, result.Error
}

// Get user by email
func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Model(&models.User{}).Where("email = ?", &email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Get user by ID
func (r *AuthRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := r.DB.Where("id = ?", id).First(&user)
	return &user, result.Error
}

// Get user by Google ID
func (r *AuthRepository) GetUserByGoogleID(googleID string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("google_id = ?", googleID).First(&user)
	return &user, result.Error
}

// Authenticate user by email and password
func (r *AuthRepository) AuthenticateUser(email, password string) (*models.User, error) {
	var user *models.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// Verify password using bcrypt
	if err := hashpasscode.CheckPassword()

	return &user, nil
}

// Create trial subscription for user
func (r *AuthRepository) CreateTrialSubscription(userID uuid.UUID) error {
	subscription := &models.Subscription{
		UserID:       userID,
		Subscription: models.PlanType_Trial,
		Status:       models.PlanStatus_Active,
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(0, 0, 7), // 7-day trial
	}

	result := r.DB.Create(subscription)
	return result.Error
}

// Get user subscription
func (r *AuthRepository) GetUserSubscription(userID uuid.UUID) (*models.Subscription, error) {
	var subscription models.Subscription
	result := r.DB.Where("user_id = ?", userID).First(&subscription)
	return &subscription, result.Error
}

// Update user status (e.g., activate/suspend)
func (r *AuthRepository) UpdateUserStatus(user *models.User) error {
	result := r.DB.Save(user)
	return result.Error
}
