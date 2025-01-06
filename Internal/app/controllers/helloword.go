package controllers

import (
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
		"message": "Hello, World!",
		"username": Username,
		"version": "0.1",
	})
}