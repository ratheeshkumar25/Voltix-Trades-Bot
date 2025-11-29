package middleware

import (
	"ratheeshkumar25/github.com/trading_bot/auth-service/pkg/http"

	"gorm.io/gorm"
)

type Middleware struct {
	//Db connection and other fields as needed
	App *http.App
	DB  *gorm.DB
}
