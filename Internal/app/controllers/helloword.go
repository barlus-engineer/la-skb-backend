package controllers

import (
	"la-skb/text"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	session := sessions.Default(c)
	Username := session.Get("username")
	SessionID := session.ID()
	if Username == nil {
		Username = "null"
	}
	if SessionID == "" {
		SessionID = "null"
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  text.Set["helloworld.helloworld"],
		"username": Username,
		"session_id": SessionID,
		"version":  text.Set["app.version"],
	})
}
