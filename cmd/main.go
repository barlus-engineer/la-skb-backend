package main

import (
	"fmt"
	"laskb-server-api/config"
	"laskb-server-api/internal/core/adapters/handler"
	"laskb-server-api/pkg/logger"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Server is starting...")
	InitRouter()
}

func InitRouter() {
	var (
		router    *gin.Engine
		log	string
		cfg       = config.LoadConfig()
		serveAddr = fmt.Sprintf("%s:%s", cfg.Postgres.Host, cfg.Server.Port)
	)

	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
		log += logger.Sinfo("HTTP server is running in production mode")
		log += "\n"
	} else {
		router = gin.Default()
		log += logger.Sinfo("HTTP server is running in development mode")
		log += "\n"
	}

	if cfg.Server.Pprof {
		pprof.Register(router)
	}

	ping := router.Group("/ping")
	{
		ping.GET("/", handler.Ping)
	}

	log += logger.Sinfof("Listening and serving HTTP on %s", serveAddr)
	log += "\n"
	fmt.Print(log)

	err := router.Run(serveAddr)
	if err != nil {
		logger.Fatalf("Error starting server: %v", err)
	}
}
