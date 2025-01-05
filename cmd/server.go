package main

import (
	"fmt"
	"la-skb/Internal/routers"
	"la-skb/config"
	"la-skb/pkg/logger"
)

func main() {
    cfg := config.LoadConfig()
    serve :=  routers.SetupServer()

    Addr := fmt.Sprintf("%s:%s", cfg.IP, cfg.Port)
    logger.Info(fmt.Sprintf("Server is started on http://%s", Addr))
    serve.Run(Addr)
}