package repo

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
)

// Create session
func (r *AuthRepository) CreateSession(session *models.Session) error {
	result := r.DB.Create(session)
	return result.Error
}
