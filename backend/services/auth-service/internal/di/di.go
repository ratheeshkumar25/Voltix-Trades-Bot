package di

import (
	"log"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/config"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/database"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/handlers"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/repo"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App) *handlers.AuthServiceHandler {
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	} // Assuming LoadConfig exists or I need to check config package

	// Connect DB
	db := database.ConnectDB(cfg)
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	// Init Repo
	authRepo := repo.NewAuthRepository(db)

	// Init Service
	userService := services.NewUserService(authRepo)

	// Init Handler
	authHandler := handlers.NewAuthServiceHandler(userService, app)

	return authHandler
}
