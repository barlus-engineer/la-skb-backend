package routers

import (
	"github.com/gin-gonic/gin"

	"la-skb/Internal/app/controllers"
)

func SetupServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	
	app := gin.New()

	app.GET("/", controllers.HelloWorld)

	return app
}