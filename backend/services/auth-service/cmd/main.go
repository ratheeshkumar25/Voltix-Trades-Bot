package main

import (
	"log"
	"os"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/api"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/di"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	// Global middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_ORIGINS"),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	// Initialize DI and get handler
	authHandler := di.Initialize(app)

	// Initialize OAuth
	authHandler.InitOAuth()

	// Setup routes
	// Setup routes
	api.SetupRoutes(app, authHandler)

	// Start server
	port := getEnv("PORT", "3001")
	log.Printf("🚀 Auth Service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   err.Error(),
		"service": "auth-service",
	})
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
