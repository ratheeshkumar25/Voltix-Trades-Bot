package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type Trade struct {
	gorm.Model
	Exchange string
	Symbol   string
	Side     string
	Price    float64
	Quantity float64
	Profit   float64
	Status   string
}

func Connect() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/trading_bot?charset=utf8mb4&parseTime=True&loc=Local"
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v. Continuing without DB for now (mock mode).", err)
		return
	}

	log.Println("Database connected")
	DB.AutoMigrate(&User{}, &Trade{})
}
