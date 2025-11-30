package handlers

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthServiceHamdler struct {
	SVC  services.UserService
	Http *fiber.App
}

func NewAuthServiceHandler(svc services.UserService, app *fiber.App) *AuthServiceHamdler {
	return &AuthServiceHamdler{SVC: svc, Http: app}
}
