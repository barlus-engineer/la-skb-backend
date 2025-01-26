package main

import (
	"laskb-server-api/config"
	"laskb-server-api/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	logger.Info("Hi i am barlus")
	logger.Infof("Hi brother %s", cfg.Server.Port)
}