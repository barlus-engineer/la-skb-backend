package controllers

import (
	"la-skb/Internal/app/appError"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/services"
	"la-skb/text"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthForm

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": text.AuthSignUpIncompleteForm,
		})
		return
	}

	if err := Auth.SignUp(); err != nil {
		switch err {
			case appError.ErrSignup: {
				c.Status(http.StatusInternalServerError)
			}
			case appError.ErrSignupUserExists: {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": text.AuthSignUpUserExists,
				})
			}
			default: {
				c.Status(http.StatusInternalServerError)
			}
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": text.AuthSignUpSuccess,
	})
}

func SignIn(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthForm

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": text.AuthSignInIncompleteForm,
		})
		return
	}
	if err := Auth.SignIn(); err != nil {
		switch err {
			case appError.ErrSignin: {
				c.Status(http.StatusInternalServerError)
			}
			case appError.ErrSigninUserNotFound: {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": text.AuthSignInUserNotFound,
				})
			}
			case appError.ErrSigninPasswordIncorrect: {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": text.AuthSignInPasswordIncorrect,
				})
			}
			default: {
				c.Status(http.StatusInternalServerError)
			}
		}
		return
	}
	token, err := Auth.AccessJWT()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": text.AuthSignInSuccess,
		"token": token,
	})
}

// func SignOut(c *gin.Context) {
// 	session := sessions.Default(c)
// 	session.Clear()
// 	session.Save()
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": text.Set["auth.signout.success"],
// 	})
// }

func DeleteAccount(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthForm

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": text.AuthDeleteAccountIncompleteForm,
		})
		return
	}

	if err := Auth.DeleteAccount(); err != nil {
		switch err {
			case appError.ErrDelAccount: {
				c.Status(http.StatusInternalServerError)
			}
			case appError.ErrDelAccountUserNotFound: {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": text.AuthSignInUserNotFound,
				})
			}
			case appError.ErrDelAccountPasswordIncorrect: {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": text.AuthSignInPasswordIncorrect,
				})
			}
			default: {
				c.Status(http.StatusBadRequest)
			}
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": text.AuthDeleteAccountSuccess,
	})
}
