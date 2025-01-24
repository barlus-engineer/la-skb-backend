package controllers

import (
	"la-skb/text"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  text.HelloWorld,
		"version":  text.AppVersion,
	})
}
