package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"la-skb/Internal/app/controllers"
	"la-skb/config"
)

func SetupServer() *gin.Engine {
	cfg := config.LoadConfig()

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	
	store := cookie.NewStore([]byte(cfg.Secret))
	app.Use(
		sessions.Sessions("username", store),
	)

	app.GET("/", controllers.HelloWorld)

	// auth
	auth := app.Group("auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/signin", controllers.SignIn)
	}

	return app
}