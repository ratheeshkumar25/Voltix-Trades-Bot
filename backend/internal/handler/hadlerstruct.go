// Developer : Ratheesh Kumar Golang Developer
package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/services"
	httpPkg "github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/http"
)

// Authservicehandler handler login and exchange login and switch
type AuthServiceHandler struct {
	SVC services.UserServiceInterface
	// App is the custom wrapper that provides helper response methods and logger
	App *httpPkg.App
	// Fiber instance for direct route registration or middleware when needed
	Http *fiber.App
}

// NewAuthServiceHandler creates a new AuthServiceHandler instance
func NewAuthServiceHandler(svc services.UserServiceInterface, app *httpPkg.App) *AuthServiceHandler {
	var fiberApp *fiber.App
	if app != nil {
		fiberApp = app.Fiber
	}
	return &AuthServiceHandler{SVC: svc, App: app, Http: fiberApp}
}
