package controllers

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	text := map[string]string {
		"message": "Hello, World!",
		"version": "0.1",
	}
	c.JSON(200, text)
}