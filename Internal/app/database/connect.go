package database

import (
	"la-skb/config"
	"la-skb/pkg/logger"
	"log"

	"gorm.io/driver/mysql"
	 GormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	cfg := config.LoadConfig()
	dsn := cfg.DbURI

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: GormLogger.Discard,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	logger.Info("Connected to MariaDB successfully!")
}

func GetDB() *gorm.DB {
	return db
}
