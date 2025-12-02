package main

import "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/di"

func main() {
	app := fiber.New()

	// Initialize dependencies and handlers
	handler := di.Initialize(app)

	if handler == nil {
		log.Fatal("failed to initialize application")
	}

	// The actual server start may be elsewhere; for now just start Fiber
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
