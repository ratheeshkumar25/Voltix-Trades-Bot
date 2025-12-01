package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("Starting Trade Automation Service...")

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Create Fiber app for WebSocket
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/trades", websocket.New(func(c *websocket.Conn) {
		log.Println("WebSocket connection established")
		for {
			// Read message (optional)
			_, _, err := c.ReadMessage()
			if err != nil {
				break
			}
			// Write message (placeholder for real-time updates)
			if err := c.WriteMessage(websocket.TextMessage, []byte("Trade update")); err != nil {
				break
			}
		}
	}))

	// Start Fiber in a goroutine
	go func() {
		log.Fatal(app.Listen(":3002"))
	}()

	// Subscribe to trade signals
	nc.Subscribe("trade.signal", func(m *nats.Msg) {
		log.Printf("Received trade signal: %s", string(m.Data))
		// Broadcast to WebSockets
		// TODO: Implement broadcast mechanism
	})

	log.Println("Trade Automation Service is running")

	// Wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
