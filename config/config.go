package config

import (
	"la-skb/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IP   string
	Port string
	Db   Db
}

type Db struct {
	Host     string
	User     string
	Password string
	Database string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Warning("Failed to load .env file: ")
		logger.Info("Fallback: Switching to system environment variables")
	}

	cfg_IP := os.Getenv("IP")
	if cfg_IP == "" {
		cfg_IP = "127.0.0.1"
	}
	cfg_Port := os.Getenv("PORT")
	if cfg_Port == "" {
		cfg_Port = "3432"
	}
	cfg_Db_Host := os.Getenv("DB_HOST")
	if cfg_Db_Host == "" {
		cfg_Db_Host = "127.0.0.1"
	}
	cfg_Db_User := os.Getenv("DB_USER")
	cfg_Db_Password := os.Getenv("DB_PASSWORD")
	cfg_Db_Database := os.Getenv("DB_DATABASE")

	cfg := Config{
		IP:   cfg_IP,
		Port: cfg_Port,
		Db: Db{
			Host:     cfg_Db_Host,
			User:     cfg_Db_User,
			Password: cfg_Db_Password,
			Database: cfg_Db_Database,
		},
	}

	return &cfg
}
