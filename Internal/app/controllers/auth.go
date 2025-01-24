package controllers

import (
	"fmt"
	"la-skb/Internal/app/appError"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/services"
	"la-skb/text"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthForm

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		handleError(c, appError.ErrBindBodyWithJSON)
		return
	}

	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		handleError(c, appError.ErrSigninSignInIncompleteForm)
		return
	}

	if err := Auth.SignUp(); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": text.SignUpSuccess,
	})
}

func SignIn(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthForm

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		handleError(c, appError.ErrBindBodyWithJSON)
		return
	}
	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		handleError(c, appError.ErrSigninSignInIncompleteForm)
		return
	}
	if err := Auth.SignIn(); err != nil {
		handleError(c, err)
		return
	}
	refreshToken, err := services.RefreshJWT(Auth)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	c.SetCookie(
        "refresh_jwt",
        refreshToken,
        math.MaxInt32,
        "/",
        cfg.IP,
        true,
        true,
    )
	c.JSON(http.StatusOK, gin.H{
		"message": text.SignInSuccess,
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
		handleError(c, appError.ErrDelIncompleteForm)
		return
	}

	if err := Auth.DeleteAccount(); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": text.DeleteAccountSuccess,
	})
}

// =================================================

func handleError(c *gin.Context, err error) {
	switch err {
		case appError.ErrBindBodyWithJSON: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.BindBodyWithJSONError,
			})
		}
		// SignUp - HandleError
		case appError.ErrSigninSignInIncompleteForm: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignUpIncompleteForm,
			})
		}
		case appError.ErrSignup: {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": text.SignUpIncompleteForm,
			})
		}
		case appError.ErrSignupUserExists: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignUpUserExists,
			})
		}
		// SignIn - HandleError
		case appError.ErrSigninSignInIncompleteForm: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInIncompleteForm,
			})
		}
		case appError.ErrSignin: {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": text.SignInServerError,
			})
		}
		case appError.ErrSigninUserNotFound: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInUserNotFound,
			})
		}
		case appError.ErrSigninPasswordIncorrect: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInPasswordIncorrect,
			})
		}
		// DeleteAccount
		case appError.ErrDelIncompleteForm: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.DeleteAccountIncompleteForm,
			})
		}
		case appError.ErrDelAccount: {
			c.Status(http.StatusInternalServerError)
		}
		case appError.ErrDelAccountUserNotFound: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInUserNotFound,
			})
		}
		case appError.ErrDelAccountPasswordIncorrect: {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInPasswordIncorrect,
			})
		}
		default: {
			c.Status(http.StatusInternalServerError)
		}
	}
}