package controllers

import (
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/services"
	"la-skb/text"
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
			"error": text.Set["auth.signup.incomplete_form"],
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
			"error": text.Set["auth.signin.incomplete_form"],
		})
		return
	}

	result := Auth.SignIn()
	if result.Status == http.StatusOK {
		session := sessions.Default(c)
		session.Set("username", data.Username)
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
		"message": text.Set["auth.signout.success"],
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
			"error": text.Set["auth.delete_account.incomplete_form"],
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
