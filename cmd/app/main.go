package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"la-skb/Internal/app/database"
	"la-skb/Internal/routers"
	"la-skb/config"
	"la-skb/pkg/logger"
)

func main() {
	// database
	database.InitDB()
	db := database.GetDB()
	database.MigrateDB(db)

	// server
	cfg := config.LoadConfig()
	serve := routers.SetupServer()

	Addr := fmt.Sprintf("%s:%s", cfg.IP, cfg.Port)
	logger.Info(fmt.Sprintf("Server is started on http://%s", Addr))

	// Signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Print("\r")
		logger.Info("Shutting down server....")
		os.Exit(0)
	}()

	// Goroutine to listen for "stop" command
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			if input == "stop\n" {
				logger.Info("Received stop command. Shutting down server....")
				os.Exit(0)
			}
		}
	}()

	serve.Run(Addr)
}
