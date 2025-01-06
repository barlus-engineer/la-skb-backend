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
	var Auth services.Auth
	var data entities.AuthFormData

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ຂໍ້ມູນບໍ່ຄົບຖ້ວນ",
		})
		return
	}

	result := Auth.SignUp()
	c.JSON(result.Status, gin.H{
		"message": result.Message,
	})
}

func SignIn(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthFormData

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ຂໍ້ມູນບໍ່ຄົບຖ້ວນ",
		})
		return
	}

	result := Auth.SignIn()
	if result.Status == http.StatusOK {
		session := sessions.Default(c)
		session.Set(Auth.Username, "{}")
		session.Save()
	}
	c.JSON(result.Status, gin.H{
		"message": result.Message,
	})
}

func SignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "ອອກຈາກລະບົບສຳເລັດແລ້ວ",
	})
}

func DeleteAccount(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthFormData

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ຂໍ້ມູນບໍ່ຄົບຖ້ວນ",
		})
		return
	}

	result := Auth.DeleteAccount()
	if result.Status == http.StatusOK {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
	}
	c.JSON(result.Status, gin.H{
		"message": result.Message,
	})
}
