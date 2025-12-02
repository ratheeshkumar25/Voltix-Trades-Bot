package main

import (
	_ "github.com/ratheeshkumar25/Voltix-Trades-Bot/docs"
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/app"
)

// @title Trading Bot API
// @version 1.0
// @description Trading Bot Backend API with OAuth2 and JWT authentication
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app.StartApp()
}
