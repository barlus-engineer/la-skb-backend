package controllers

import (
	"la-skb/Internal/app/text"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	session := sessions.Default(c)
	Username := session.Get("username")
	if Username == nil {
		Username = "null"
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  text.Set["helloworld.helloworld"],
		"username": Username,
		"version":  text.Set["app.version"],
	})
}
