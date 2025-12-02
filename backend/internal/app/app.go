package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/config"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/di"
)

func StartApp() {

	// Create fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Voltix-Trades-Bot",
		ServerHeader: "Voltix-Trades-Bot",
	})

	//Global middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_ORIGINS"),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	authHandler := di.Initialize(app)

	authHandler.InitOAuth()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Start server
	port := cfg.VOLTIX_PORT
	log.Printf("🚀 Auth Service running on port %s", port)
	log.Fatal(app.Listen(":" + port))

}
