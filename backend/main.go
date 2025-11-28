package main

import (
	"log"
	"trading_bot/internal/api"
	"trading_bot/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to DB
	db.Connect()

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("Error connecting to NATS: %v", err)
	} else {
		log.Println("Connected to NATS")
		defer nc.Close()
	}

	// Setup Fiber
	app := fiber.New()
	api.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
