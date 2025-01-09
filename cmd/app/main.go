package main

import (
	"fmt"
	"la-skb/Internal/app/database"
	"la-skb/Internal/routers"
	"la-skb/config"
	"la-skb/pkg/logger"
	"la-skb/text"
)

func main() {
	// database
	database.InitDB()
	db := database.GetDB()
	database.MigrateDB(db)

	// lang
	text.InitLang()

	// server
	cfg := config.LoadConfig()
	serve := routers.SetupServer()

	Addr := fmt.Sprintf("%s:%s", cfg.IP, cfg.Port)
	logger.Info(fmt.Sprintf("Server is started on http://%s", Addr))
	serve.Run(Addr)
}
