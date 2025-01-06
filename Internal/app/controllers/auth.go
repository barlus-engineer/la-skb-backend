package controllers

import (
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/services"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var data entities.AuthFormData
	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Username := data.Username
	Password := data.Password

	if strings.TrimSpace(Username) == "" || strings.TrimSpace(Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ຂໍ້ມູນບໍ່ຄົບຖ້ວນ",
		})
		return
	}

	result := services.SignUp(Username, Password)
	c.JSON(result.Status, gin.H{
		"message": result.Message,
	})
}

func SignIn(c *gin.Context) {
	var data entities.AuthFormData
	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Username := data.Username
	Password := data.Password

	if strings.TrimSpace(Username) == "" || strings.TrimSpace(Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ຂໍ້ມູນບໍ່ຄົບຖ້ວນ",
		})
		return
	}

	result := services.SignIn(Username, Password)
	if result.Status == http.StatusOK {
		session := sessions.Default(c)
		session.Set("username", Username)
		session.Save()
	}
	c.JSON(result.Status, gin.H{
		"message": result.Message,
	})
}
