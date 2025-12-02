package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/handler"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/middleware"
)

func SetupRoutes(app *fiber.App, h *handler.AuthServiceHandler) {
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", h.EmailLoginHandler)
	auth.Post("/register", h.EmailRegisterHandler)  // Frontend uses this endpoint
	auth.Post("/google", h.GoogleCredentialHandler) // Credential-based OAuth for frontend
	auth.Get("/google/login", h.GoogleLoginHandler) // Redirect-based OAuth
	auth.Get("/google/callback", h.GoogleCallbackHandler)

	// User/account routes - currently handlers implemented in handler package
	user := api.Group("/v1/user", middleware.AuthMiddleware)
	// NOTE: handlers `UpdateUser` and `DeleteUser` were not present in handler package
	// Register only implemented account handlers
	user.Get("/me", h.MeHandler)
	user.Post("/account", h.EmailRegisterHandler)
	user.Put("/account", h.UpdateAccountHandler)
	user.Delete("/account", h.DeleteAccountHandler)

	// Exchange routes
	exchange := api.Group("/v1/exchange", middleware.AuthMiddleware)
	exchange.Post("/", h.AddExchangeCredential)
	exchange.Get("/", h.GetExchangeCredentials)
	exchange.Post("/switch", h.SwitchExchangeAccount)
}
