package di

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/config"
	dbpkg "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/db"
	handlers "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/handler"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/repo"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/services"
	httpPkg "github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/http"
	loggerPkg "github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/logger"
)

// Initialize wires dependencies and returns an auth handler instance
func Initialize(app *fiber.App) *handlers.AuthServiceHandler {
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Connect DB
	db := dbpkg.ConnectDB(cfg)
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	// Init Repo
	authRepo := repo.NewAuthRepository(db)

	// Init Logger and http App wrapper
	logInst := loggerPkg.GetLogger()
	httpApp := httpPkg.NewApp(app, logInst)

	// Init Service (note: services.NewUserService expects *http.App, repo.AuthRepositoryInterface)
	userService := services.NewUserService(httpApp, authRepo)

	// Init Handler
	authHandler := handlers.NewAuthServiceHandler(userService, httpApp)

	return authHandler
}
