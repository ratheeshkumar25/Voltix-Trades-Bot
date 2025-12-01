package handlers

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOAuthConfig *oauth2.Config
)

// InitOAuth initializes Google OAuth configuration
func (h *AuthServiceHandler) InitOAuth() {
	googleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:3001/api/auth/google/callback"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

// GoogleLoginHandler initiates Google OAuth flow
func (h *AuthServiceHandler) GoogleLoginHandler(c *fiber.Ctx) error {
	state, err := services.GenerateSecureToken(32)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate state"})
	}

	// Store state in session/cookie for validation
	c.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HTTPOnly: true,
		Secure:   true,
		MaxAge:   300, // 5 minutes
	})

	url := googleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.JSON(fiber.Map{"url": url})
}

// GoogleCallbackHandler handles Google OAuth callback
func (h *AuthServiceHandler) GoogleCallbackHandler(c *fiber.Ctx) error {
	// Validate state
	state := c.Query("state")
	storedState := c.Cookies("oauth_state")
	if state != storedState {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid state parameter"})
	}

	// Exchange code for token
	code := c.Query("code")
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	// Get user info from Google
	client := googleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user info"})
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var googleUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	json.Unmarshal(data, &googleUser)

	// Check if user exists
	user, err := h.SVC.GetUserByGoogleID(googleUser.ID)
	if err != nil {
		// Create new user
		user, err = h.SVC.CreateUser(googleUser.Email, "", &googleUser.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
		}
	}

	// Generate JWT
	jwtToken, err := services.GenerateJWT(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Create session
	session := &models.Session{
		UserID:    user.ID,
		Token:     jwtToken,
		ExpiresAt: token.Expiry,
	}
	if err := h.SVC.CreateSession(session); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create session"})
	}

	return c.JSON(fiber.Map{
		"token": jwtToken,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

// EmailLoginHandler handles email/password login
func (h *AuthServiceHandler) EmailLoginHandler(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Authenticate user
	user, err := h.SVC.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT
	jwtToken, err := services.GenerateJWT(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Create session
	session := &models.Session{
		UserID:    user.ID,
		Token:     jwtToken,
		ExpiresAt: time.Now().Add(24 * time.Hour), // Match JWT expiration
	}
	if err := h.SVC.CreateSession(session); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create session"})
	}

	return c.JSON(fiber.Map{
		"token": jwtToken,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

// RegisterHandler handles user registration
func (h *AuthServiceHandler) RegisterHandler(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Create user with trial
	user, err := h.SVC.CreateUser(req.Email, req.Password, nil)
	if err != nil {
		if err == services.ErrEmailExists {
			return c.Status(409).JSON(fiber.Map{"error": "Email already exists"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Generate JWT
	jwtToken, err := services.GenerateJWT(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(201).JSON(fiber.Map{
		"token": jwtToken,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
		"message": "Account created with 7-day free trial",
	})
}

// MeHandler returns current user info
func (h *AuthServiceHandler) MeHandler(c *fiber.Ctx) error {
	// Get user from context (set by auth middleware)
	user := c.Locals("user").(*models.User)

	// Get subscription
	subscription, err := h.SVC.GetUserSubscription(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get subscription"})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"subscription": fiber.Map{
			"plan":           subscription.Subscription,
			"status":         subscription.Status,
			"days_remaining": subscription.DaysRemaining(),
			"end_date":       subscription.EndDate,
		},
	})
}

// AddExchangeCredential adds a new exchange credential
func (h *AuthServiceHandler) AddExchangeCredential(c *fiber.Ctx) error {
	type ExchangeRequest struct {
		ExchangeType models.ExchangeType `json:"exchange_type"`
		APIKey       string              `json:"api_key"`
		APISecret    string              `json:"api_secret"`
	}
	var req ExchangeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user := c.Locals("user").(*models.User)
	cred, err := h.SVC.AddExchangeCredential(user.ID, req.ExchangeType, req.APIKey, req.APISecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add exchange credential"})
	}

	return c.Status(201).JSON(cred)
}

// GetExchangeCredentials gets all exchange credentials for a user
func (h *AuthServiceHandler) GetExchangeCredentials(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	creds, err := h.SVC.GetExchangeCredentials(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get exchange credentials"})
	}
	return c.JSON(creds)
}

// SwitchExchangeAccount switches the active exchange account
func (h *AuthServiceHandler) SwitchExchangeAccount(c *fiber.Ctx) error {
	type SwitchRequest struct {
		CredentialID string `json:"credential_id"`
	}
	var req SwitchRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	credentialID, err := uuid.Parse(req.CredentialID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credential ID"})
	}

	user := c.Locals("user").(*models.User)
	if err := h.SVC.SwitchExchangeAccount(user.ID, credentialID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to switch exchange account"})
	}

	return c.JSON(fiber.Map{"message": "Exchange account switched successfully"})
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
