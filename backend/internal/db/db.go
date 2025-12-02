package db

import (
	"fmt"
	"log"

	"github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/config"
	models "github.com/ratheeshkumar25/Voltix-Trades-Bot/internal/modles"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

func ConnectDB(config *config.Config) *gorm.DB {
	host := config.DBHost
	port := config.DBPort
	user := config.DBUser
	password := config.DBPassword
	dbname := config.DBName
	sslmode := config.DBSSLMode

	//
	log.Printf("Connecting to DB: host=%s, user=%s, password=%s, dbname=%s, port=%s, sslmode=%s\n", host, user, password, dbname, port, sslmode)

	//construct the DSN connection correctly for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Subscription{},
		&models.Notification{},
		&models.Session{},
	)

	if err != nil {
		log.Printf("error while migrating %v", err.Error())
		return nil
	}
	DB = db // Assign to global variable
	return db

}
