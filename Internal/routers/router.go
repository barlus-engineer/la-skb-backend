package routers

import (
	"github.com/gin-contrib/sessions"
	gormstore "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"

	"la-skb/Internal/app/controllers"
	"la-skb/Internal/app/database"
	"la-skb/config"
)

func SetupServer() *gin.Engine {
	cfg := config.LoadConfig()

	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	db := database.GetDB()
	
	store := gormstore.NewStore(db, true, []byte(cfg.Secret))
	// Fix it, save only sessoin key
	app.Use(
		sessions.Sessions("username", store),
	)

	app.GET("/", controllers.HelloWorld)

	auth := app.Group("auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/signin", controllers.SignIn)
		auth.POST("/signout", controllers.SignOut)
		auth.POST("/delete_account", controllers.DeleteAccount)
	}

	return app
}