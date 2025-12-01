package handlers

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthServiceHandler struct {
	SVC  services.UserServiceInterface
	Http *fiber.App
}

func NewAuthServiceHandler(svc services.UserServiceInterface, app *fiber.App) *AuthServiceHandler {
	return &AuthServiceHandler{SVC: svc, Http: app}
}
