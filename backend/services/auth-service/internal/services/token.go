package services

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateSecureToken generates a random token
func GenerateSecureToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}
