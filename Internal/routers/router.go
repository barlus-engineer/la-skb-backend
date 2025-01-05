package routers

import (
	"github.com/gin-gonic/gin"

	"la-skb/Internal/app/controllers"
)

func SetupServer() *gin.Engine {
	app := gin.New()

	app.GET("/", controllers.HelloWorld)

	return app
}