package middleware

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/database"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"
	"ratheeshkumar25/github.com/trading_bot/auth-service/pkg/oauth"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AuthMiddleware validates JWT token and loads user
func AuthMiddleware(c *fiber.Ctx) error {
	// Get token from Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing authorization header"})
	}

	// Extract token from "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid authorization format"})
	}

	tokenString := parts[1]

	// Validate JWT
	claims, err := oauth.ValidateJWT(tokenString)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Load user from database
	var user models.User
	if err := database.DB.First(&user, "id = ?", claims.UserID).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	// Check if user is active
	if !user.IsActive {
		return c.Status(403).JSON(fiber.Map{"error": "User account is suspended"})
	}

	// Store user in context
	c.Locals("user", &user)
	c.Locals("userID", user.ID)

	return c.Next()
}

// AdminMiddleware checks if user has admin role
func AdminMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if user.Role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "Admin access required"})
	}
	return c.Next()
}

// SubscriptionMiddleware checks if user has active subscription
func SubscriptionMiddleware(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var subscription models.Subscription
	if err := database.DB.Where("user_id = ?", userID).First(&subscription).Error; err != nil {
		return c.Status(403).JSON(fiber.Map{
			"error":   "No active subscription",
			"message": "Please subscribe to continue using trading features",
		})
	}

	if !subscription.IsActive() {
		return c.Status(403).JSON(fiber.Map{
			"error":          "Subscription expired",
			"message":        "Your trial/subscription has expired. Please upgrade to continue.",
			"plan":           subscription.Subscription,
			"expired_at":     subscription.EndDate,
			"days_remaining": subscription.DaysRemaining(),
		})
	}

	// Store subscription in context
	c.Locals("subscription", &subscription)

	return c.Next()
}
