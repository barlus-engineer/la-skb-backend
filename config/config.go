package config

import (
	"la-skb/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	IP   			string
	Port 			string
	Secret			string
	PublicSecret	string
	DbURI			string
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
	cfg_Db_Secret := os.Getenv("SECRET")
	cfg_Db_PublicSecret := os.Getenv("SECRET")
	cfg_Db_URI := os.Getenv("DB_URI")

	cfg := Config{
		IP:   cfg_IP,
		Port: cfg_Port,
		Secret: cfg_Db_Secret,
		PublicSecret: cfg_Db_PublicSecret,
		DbURI: cfg_Db_URI,
	}

	return &cfg
}
