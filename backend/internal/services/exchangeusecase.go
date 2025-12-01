package services

import (
	"errors"

	"github.com/google/uuid"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
)

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
