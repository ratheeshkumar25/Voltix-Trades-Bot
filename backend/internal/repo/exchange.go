package repo

import (
	"github.com/google/uuid"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
)

// AddExchangeCredential adds a new exchange credential
func (r *AuthRepository) AddExchangeCredential(cred *models.ExchangeCredential) error {
	result := r.DB.Create(cred)
	return result.Error
}

// GetExchangeCredentials gets all exchange credentials for a user
func (r *AuthRepository) GetExchangeCredentials(userID uuid.UUID) ([]models.ExchangeCredential, error) {
	var creds []models.ExchangeCredential
	result := r.DB.Where("user_id = ?", userID).Find(&creds)
	return creds, result.Error
}

// UpdateExchangeCredential updates an exchange credential
func (r *AuthRepository) UpdateExchangeCredential(cred *models.ExchangeCredential) error {
	result := r.DB.Save(cred)
	return result.Error
}
