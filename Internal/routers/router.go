package routers

import (
	"github.com/gin-gonic/gin"

	"la-skb/Internal/app/controllers"
)

func SetupServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	app.GET("/", controllers.HelloWorld)

	auth := app.Group("auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/signin", controllers.SignIn)
		// auth.POST("/signout", controllers.SignOut)
		auth.POST("/delete_account", controllers.DeleteAccount)
	}

	return app
}