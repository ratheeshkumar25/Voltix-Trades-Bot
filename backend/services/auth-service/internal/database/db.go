package database

import (
	"fmt"
	"log"

	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/config"
	"ratheeshkumar25/github.com/trading_bot/auth-service/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global database connection

func ConnectDB(config *config.Config) *gorm.DB {
	host := config.DBHost
	port := config.DBPort
	user := config.DBUser
	password := config.DBPassword
	dbname := config.DBName
	sslmode := config.DBSSLMode

	//
	log.Printf("Connecting to DB: host=%s, user=%s, password=%s, dbname=%s, port=%s, sslmode=%s\n", host, user, password, dbname, port, sslmode)

	//construct the DSN connection correctly
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC", host, user, password, dbname, port, sslmode)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Subscription{},
		&models.Notification{},
		&models.Session{},
	)

	if err != nil {
		log.Printf("error while migrating %v", err.Error())
		return nil
	}
	return DB

}
