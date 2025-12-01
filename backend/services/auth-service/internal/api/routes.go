package api

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/handlers"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handlers.AuthServiceHandler) {
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", h.EmailLoginHandler)
	auth.Get("/google", h.GoogleLoginHandler)
	auth.Get("/google/callback", h.GoogleCallbackHandler)

	// User routes
	user := api.Group("/v1/user", middleware.AuthMiddleware)
	user.Post("/", h.CreateUser)
	user.Get("/:email", h.GetUserByEmail) // Changed to param for better REST practice, though handler expects body currently?
	// Wait, the handler GetUserByEmail uses c.Params("email") now after my fix.
	// But let's check the handler again to be sure.
	user.Get("/id/:id", h.GetUserByID)
	user.Get("/google/:google_id", h.GetUserByGoogleID)
	user.Put("/", h.UpdateUser)
	user.Delete("/", h.DeleteUser)

	// Exchange routes
	exchange := api.Group("/v1/exchange", middleware.AuthMiddleware)
	exchange.Post("/", h.AddExchangeCredential)
	exchange.Get("/", h.GetExchangeCredentials)
	exchange.Post("/switch", h.SwitchExchangeAccount)
}
