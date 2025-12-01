package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Println("Starting Admin & Payment Service...")

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")

	// Admin routes
	admin := api.Group("/admin")
	admin.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "List of users"})
	})

	// Payment routes
	payment := api.Group("/payment")
	payment.Post("/process", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Payment processed"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	log.Printf("🚀 Admin-Payment Service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
