package config

import (
	"la-skb/pkg/logger"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	IP   			string
	Port 			string
	Secret			string
	PublicSecret	string
	CacheURI		string
	CacheTime		time.Duration
	DbURI			string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Warning("Failed to load .env file: ")
		logger.Info("Fallback: Switching to system environment variables")
	}

	cfgIP := os.Getenv("IP")
	if cfgIP == "" {
		cfgIP = "127.0.0.1"
	}
	cfgPort := os.Getenv("PORT")
	if cfgPort == "" {
		cfgPort = "3432"
	}

	cfgDbSecret := os.Getenv("SECRET")
	cfgDbPublicSecret := os.Getenv("PUBLIC_SECRET")

	cfgDbURI := os.Getenv("DB_URI")

	cfgCacheURI := os.Getenv("CACHE_URI")
	cfgCacheTime := 30 * time.Minute
	if t := os.Getenv("CACHE_TIME"); t != "" {
		if n, err := strconv.Atoi(t); err == nil {
			cfgCacheTime = time.Duration(n) * time.Minute
		} else {
			logger.Warning("Config: Invalid CACHE_TIME value, using default!. Error: ", err)
		}
	}

	cfg := Config{
		IP:   cfgIP,
		Port: cfgPort,
		Secret: cfgDbSecret,
		PublicSecret: cfgDbPublicSecret,
		CacheURI: cfgCacheURI,
		CacheTime: cfgCacheTime,
		DbURI: cfgDbURI,
	}

	return &cfg
}
