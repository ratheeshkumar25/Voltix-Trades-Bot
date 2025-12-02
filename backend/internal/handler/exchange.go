package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
)

// ExchangeAddRequest represents payload to add exchange credentials
type ExchangeAddRequest struct {
	ExchangeType string `json:"exchange_type"`
	APIKey       string `json:"api_key"`
	APISecret    string `json:"api_secret"`
}

// @Id AddExchange
// @Summary Add exchange credential
// @Tags Exchange
// @Accept json
// @Produce json
// @Param body body handler.ExchangeAddRequest true "Add exchange credential"
// @Success 201 {object} http.HttpResponse
// @Failure 400 {object} http.HttpResponse
// @Failure 401 {object} http.HttpResponse
// @Failure 500 {object} http.HttpResponse
// @Security BearerAuth
// @Router /api/v1/exchange [post]
// AddExchangeCredential handles adding a new exchange credential for the authenticated user
func (h *AuthServiceHandler) AddExchangeCredential(c *fiber.Ctx) error {
	var req ExchangeAddRequest
	if err := c.BodyParser(&req); err != nil {
		return h.App.HttpResponseBadRequest(c, err)
	}

	// Get user ID from middleware
	userIDVal := c.Locals("userID")
	if userIDVal == nil {
		return h.App.HttpResponseUnauthorized(c, fiber.ErrUnauthorized)
	}

	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		return h.App.HttpResponseInternalServerErrorRequest(c, fiber.ErrInternalServerError)
	}

	cred, err := h.SVC.AddExchangeCredential(userID, models.ExchangeType(req.ExchangeType), req.APIKey, req.APISecret)
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return h.App.HttpResponseCreated(c, cred)
}

// GetExchangeCredentials returns all exchange credentials for the authenticated user
// @Id ListExchange
// @Summary List exchange credentials
// @Tags Exchange
// @Produce json
// @Success 200 {object} http.HttpResponse
// @Failure 401 {object} http.HttpResponse
// @Failure 500 {object} http.HttpResponse
// @Security BearerAuth
// @Router /api/v1/exchange [get]
func (h *AuthServiceHandler) GetExchangeCredentials(c *fiber.Ctx) error {
	userIDVal := c.Locals("userID")
	if userIDVal == nil {
		return h.App.HttpResponseUnauthorized(c, fiber.ErrUnauthorized)
	}
	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		return h.App.HttpResponseInternalServerErrorRequest(c, fiber.ErrInternalServerError)
	}

	creds, err := h.SVC.GetExchangeCredentials(userID)
	if err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return h.App.HttpResponseOK(c, creds)
}

// SwitchExchangeRequest is payload to switch active exchange credential
type SwitchExchangeRequest struct {
	CredentialID string `json:"credential_id"`
}

// @Id SwitchExchange
// @Summary Switch active exchange credential
// @Tags Exchange
// @Accept json
// @Produce json
// @Param body body handler.SwitchExchangeRequest true "Switch credential"
// @Success 200 {object} http.HttpResponse
// @Failure 400 {object} http.HttpResponse
// @Failure 401 {object} http.HttpResponse
// @Failure 500 {object} http.HttpResponse
// @Security BearerAuth
// @Router /api/v1/exchange/switch [post]
// SwitchExchangeAccount switches which exchange credential is active
func (h *AuthServiceHandler) SwitchExchangeAccount(c *fiber.Ctx) error {
	var req SwitchExchangeRequest
	if err := c.BodyParser(&req); err != nil {
		return h.App.HttpResponseBadRequest(c, err)
	}

	credUUID, err := uuid.Parse(req.CredentialID)
	if err != nil {
		return h.App.HttpResponseBadRequest(c, err)
	}

	userIDVal := c.Locals("userID")
	if userIDVal == nil {
		return h.App.HttpResponseUnauthorized(c, fiber.ErrUnauthorized)
	}
	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		return h.App.HttpResponseInternalServerErrorRequest(c, fiber.ErrInternalServerError)
	}

	if err := h.SVC.SwitchExchangeAccount(userID, credUUID); err != nil {
		return h.App.HttpResponseInternalServerErrorRequest(c, err)
	}

	return h.App.HttpResponseOK(c, map[string]string{"message": "switched"})
}
