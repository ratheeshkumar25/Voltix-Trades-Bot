// Developer : Ratheesh Kumar Golang Developer
package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/config"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/oauth"
	"golang.org/x/oauth2"
)

var googleOAuthConfig *oauth2.Config

// InitOAuth initializes the Google OAuth configuration
func InitOAuth(h *AuthServiceHandler) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	googleOAuthConfig = &oauth2.Config{
		ClientID:     cfg.GOOGLE_CLIENT_ID,
		ClientSecret: cfg.GOOGLE_CLIENT_SECRET,
		RedirectURL:  cfg.GOOGLE_REDIRECT_URL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.id"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
}

// @Id Create Google OAuth Handler
// @Summary Create Google OAuth Handler
// @Tags Auth
// @Produce json
// @Success 200 {object} http.HttpResponse
// @Router /api/auth/google/login [get]
// GoogleLoginHandler initiates Google OAuth flow
func (h *AuthServiceHandler) GoogleLoginHandler(c *fiber.Ctx) error {
	var user models.User
	state, err := oauth.GenerateJWT(&user)
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	// Store state in session/cookie for validation
	c.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HTTPOnly: true,
		Secure:   true,
		MaxAge:   300, // 5 minutes
	})

	if googleOAuthConfig == nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, errors.New("oauth config not initialized"))
	}

	url := googleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	// Use the http helper that writes the Created response directly
	return h.App.HttpResponseCreated(c, map[string]string{"url": url})
}

// @Id Google OAuth Callback Handler
// @Summary Google OAuth Callback Handler
// @Tags Auth
// @Produce json
// @Param code query string true "Authorization Code"
// @Param state query string true "OAuth State"
// @Success 200 {object} http.HttpResponse
// @Router /api/auth/google/callback [get]
// GoogleCallbackHandler handles Google OAuth callback
func (h *AuthServiceHandler) GoogleCallbackHandler(c *fiber.Ctx) error {
	state := c.Query("state")
	storedState := c.Cookies("oauth_state")

	// Validate state
	if state != storedState {
		return h.App.HttpResponseUnauthorized(c, errors.New("invalid oauth state"))
	}

	// Exchange code for token
	code := c.Query("code")
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	// Get user info from Google
	client := googleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var googleUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	json.Unmarshal(data, &googleUser)

	// Check if user exists or create new user
	user, err := h.SVC.GetUserByGoogleID(googleUser.ID)
	if err != nil {
		user, err = h.SVC.CreateUser(googleUser.Email, "", &googleUser.ID)
		if err != nil {
			return h.App.HttpResponseInternalServerErrorRequest(c, err)
		}
	}

	// Generate JWT for the user
	jwtToken, err := oauth.GenerateJWT(user)
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	users := models.User{
		ID:       user.ID,
		Email:    user.Email,
		Role:     user.Role,
		GoogleID: googleUser.ID,
	}

	// Create Session
	session := models.Session{
		UserID:    user.ID,
		Token:     jwtToken,
		ExpiresAt: token.Expiry,
	}

	if err := h.SVC.CreateSession(&session); err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return h.App.HttpResponseCreated(c, map[string]interface{}{
		"user":  users,
		"token": jwtToken,
	})

}

type MeResponse struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	Subscription struct {
		Plan          string `json:"plan"`
		Status        string `json:"status"`
		DaysRemaining int    `json:"days_remaining"`
		EndDate       string `json:"end_date"`
	} `json:"subscription"`
}

// @Id GetCurrentUser
// @Summary Get current user profile
// @Description Get current authenticated user's profile and subscription information
// @Tags User
// @Produce json
// @Success 200 {object} http.HttpResponse
// @Failure 401 {object} http.HttpResponse
// @Failure 500 {object} http.HttpResponse
// @Security BearerAuth
// @Router /api/v1/user/me [get]
// MeHandler returns current user info
func (h *AuthServiceHandler) MeHandler(c *fiber.Ctx) error {
	// Get user from context (set by auth middleware)
	user := c.Locals("user").(*models.User)

	// Get subscription
	subscription, err := h.SVC.GetUserSubscription(user.ID)
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	response := MeResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Role:  string(user.Role),
		Subscription: struct {
			Plan          string `json:"plan"`
			Status        string `json:"status"`
			DaysRemaining int    `json:"days_remaining"`
			EndDate       string `json:"end_date"`
		}{
			Plan:          models.PlanType_name[int32(subscription.Subscription)],
			Status:        models.PlanStatus_name[int32(subscription.Status)],
			DaysRemaining: subscription.DaysRemaining(),
			EndDate:       subscription.EndDate.String(),
		},
	}

	return h.App.HttpResponseOK(c, response)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Id Email Login Handler
// @Summary Email Login Handler
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body handler.LoginRequest true "Login Request"
// @Success 200 {object} http.HttpResponse
// @Router /api/auth/login [post]
// EmailLoginHandler handles email/password login
func (h *AuthServiceHandler) EmailLoginHandler(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return h.App.HttpResponseBadRequest(c, err)
	}

	//authenticate user
	user, err := h.SVC.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return h.App.HttpResponseUnauthorized(c, err)
	}

	h.App.Log.Logger.Infof("User logged in: %v", user.ID)

	// Generate JWT for the user
	jwtToken, err := oauth.GenerateJWT(user)
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	users := models.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	// Create Session
	session := models.Session{
		UserID:    user.ID,
		Token:     jwtToken,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := h.SVC.CreateSession(&session); err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return h.App.HttpResponseCreated(c, map[string]interface{}{
		"user":  users,
		"token": jwtToken,
	})
}

type RegisterLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Id Email Registration Handler
// @Summary Email Registration Handler
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body handler.RegisterLogin true "Register Request"
// @Success 200 {object} http.HttpResponse
// @Router /api/v1/user/account [post]
// EmailRegisterHandler handles email/password registration
func (h *AuthServiceHandler) EmailRegisterHandler(c *fiber.Ctx) error {
	var req RegisterLogin
	if err := c.BodyParser(&req); err != nil {
		return h.App.HttpResponseBadRequest(c, err)
	}

	// Create new user , before creating check if user eamil already exists
	user, err := h.SVC.GetUserByEmail(req.Email)
	if err != nil {
		user, err = h.SVC.CreateUser(req.Email, req.Password, nil)
		if err != nil {
			return h.App.HttpResponseInternalServerErrorRequest(c, err)
		}
	}
	h.App.Log.Logger.Infof("New user registered: %v", user.ID)
	users := models.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	return h.App.HttpResponseOK(c, users)
}

type UptAccount struct {
	Name  string `json:"name"`
	Email string `json:"details"`
}

// @Id				UpdateAccount
// @Summary		Update Account
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Success		200	{object}	http.HttpResponse
// @Failure		500	{object}	http.HttpResponse
// @Security		BearerAuth
// @Param			body		body	UptAccount	true	"Account Request Body"
// @Router			/api/v1/user/account [put]
func (h *AuthServiceHandler) UpdateAccountHandler(c *fiber.Ctx) error {
	// Get account id
	id := c.Params("id")
	if id == "" {
		return h.App.HttpResponseBadQueryParams(c, fmt.Errorf("id %s", errors.ErrUnsupported))
	}
	// Parse request body
	data := &UptAccount{}
	err := c.BodyParser(data)
	if err != nil {
		return h.App.HttpResponseBadRequest(c, err)
	}
	// Update user
	res, err := h.SVC.UpdateUser(&models.User{
		Name:  data.Name,
		Email: data.Email,
	})
	h.App.Log.Logger.Infof("Account updated: %v", res.ID)
	//If error return internal server error
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}
	//Return after successful updation
	return h.App.HttpResponseOK(c, res)
}

// @Id       DeleteAccount
// @Summary   Delete Account
// @Tags      Accounts
// @Produce   json
// @Success   200 {object} http.HttpResponse
// @Failure   500 {object} http.HttpResponse
// @Security  BearerAuth
// @Router    /api/v1/user/account [delete]
func (h *AuthServiceHandler) DeleteAccountHandler(c *fiber.Ctx) error {
	// Get account id
	id := c.Params("id")
	if id == "" {
		return h.App.HttpResponseBadQueryParams(c, fmt.Errorf("id %s", errors.ErrUnsupported))
	}

	account, err := h.SVC.DeleteUser(&models.User{})
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}
	h.App.Log.Logger.Infof("Account deleted: %v", account.ID)
	return h.App.HttpResponseOK(c, map[string]string{"message": "account deleted successfully"})
}
