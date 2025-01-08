package database

import (
	"la-skb/Internal/app/models"
	"la-skb/pkg/logger"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	logger.Info("Database migrated successfully!")
}