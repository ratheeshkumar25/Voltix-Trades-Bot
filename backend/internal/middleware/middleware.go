package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/http"
	"gorm.io/gorm"
)

type Middleware struct {
	//Db connection and other fields as needed
	App *http.App
	DB  *gorm.DB
}

// Default is a package-level middleware instance. It should be initialized by DI.
var Default *Middleware

// AuthMiddleware is a wrapper that delegates to the Default middleware instance.
// If Default is not initialized, it returns a 500 response indicating misconfiguration.
func AuthMiddleware(c *fiber.Ctx) error {
	if Default == nil {
		return c.Status(500).JSON(fiber.Map{"error": "middleware not initialized"})
	}
	return Default.AuthMiddleware(c)
}

// AdminMiddleware wrapper
func AdminMiddleware(c *fiber.Ctx) error {
	if Default == nil {
		return c.Status(500).JSON(fiber.Map{"error": "middleware not initialized"})
	}
	return Default.AdminMiddleware(c)
}

// SubscriptionMiddleware wrapper
func SubscriptionMiddleware(c *fiber.Ctx) error {
	if Default == nil {
		return c.Status(500).JSON(fiber.Map{"error": "middleware not initialized"})
	}
	return Default.SubscriptionMiddleware(c)
}
