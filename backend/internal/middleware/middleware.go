package middleware

import (
	"github.com/ratheeshkumar25/Voltix-Trades-Bot/pkg/http"
	"gorm.io/gorm"
)

type Middleware struct {
	//Db connection and other fields as needed
	App *http.App
	DB  *gorm.DB
}
