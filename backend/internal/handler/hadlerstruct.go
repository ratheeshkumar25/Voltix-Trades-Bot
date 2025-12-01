package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/services"
)

type AuthServiceHandler struct {
	SVC  services.UserServiceInterface
	Http *fiber.App
}

func NewAuthServiceHandler(svc services.UserServiceInterface, app *fiber.App) *AuthServiceHandler {
	return &AuthServiceHandler{SVC: svc, Http: app}
}
