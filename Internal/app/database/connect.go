package database

import (
	"la-skb/config"
	"la-skb/pkg/logger"
	"log"

	"gorm.io/driver/postgres"
	GormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	cfg := config.LoadConfig()
	dsn := cfg.DbURI

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: GormLogger.Discard,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	logger.Info("Connected to database successfully!")
}

func GetDB() *gorm.DB {
	return db
}
