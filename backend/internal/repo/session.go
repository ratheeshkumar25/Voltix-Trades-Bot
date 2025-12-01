package repo

import models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"

// Create session
func (r *AuthRepository) CreateSession(session *models.Session) error {
	result := r.DB.Create(session)
	return result.Error
}
