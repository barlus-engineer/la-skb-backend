package db

import (
	"fmt"
	"la-skb/config"
	"la-skb/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	cfg := config.LoadConfig()
	dsn := cfg.DbURI

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Alert(fmt.Sprintf("Failed to connect to the database: %v", err))
		panic(fmt.Sprintf("Cannot proceed without a database connection: %v", err)) // Optional but ensures failure is handled.
	}

	logger.Info("Connected to MariaDB successfully")
}

func GetDB() *gorm.DB {
	return db
}
