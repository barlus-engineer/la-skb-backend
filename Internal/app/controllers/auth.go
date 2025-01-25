package controllers

import (
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/services"
	"la-skb/pkg"
	"la-skb/text"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var Auth services.Auth
	var data entities.AuthForm

	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		handleError(c, pkg.ErrBindBodyWithJSON)
		return
	}

	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		handleError(c, pkg.ErrSigninIncompleteForm)
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
		handleError(c, pkg.ErrBindBodyWithJSON)
		return
	}
	Auth.Username, Auth.Password = data.Username, data.Password

	if strings.TrimSpace(Auth.Username) == "" || strings.TrimSpace(Auth.Password) == "" {
		handleError(c, pkg.ErrSigninIncompleteForm)
		return
	}
	if err := Auth.SignIn(); err != nil {
		handleError(c, err)
		return
	}
	if err := services.RefreshJWT(c, Auth); err != nil {
		handleError(c, err)
		return
	}
	accessToken, err := services.AccessJWT(Auth)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": text.SignInSuccess,
		"token":   accessToken,
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
		handleError(c, pkg.ErrDelIncompleteForm)
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
	case pkg.ErrBindBodyWithJSON:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.BindBodyWithJSONError,
			})
		}
	// SignUp - HandleError
	case pkg.ErrSigninIncompleteForm:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignUpIncompleteForm,
			})
		}
	case pkg.ErrSignup:
		{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": text.SignUpIncompleteForm,
			})
		}
	case pkg.ErrSignupUserExists:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignUpUserExists,
			})
		}
	// SignIn - HandleError
	case pkg.ErrSigninIncompleteForm:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInIncompleteForm,
			})
		}
	case pkg.ErrSignin:
		{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": text.SignInServerError,
			})
		}
	case pkg.ErrSigninUserNotFound:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInUserNotFound,
			})
		}
	case pkg.ErrSigninPasswordIncorrect:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInPasswordIncorrect,
			})
		}
	// DeleteAccount
	case pkg.ErrDelIncompleteForm:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.DeleteAccountIncompleteForm,
			})
		}
	case pkg.ErrDelAccount:
		{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": text.DeleteAccountServerError,
			})
		}
	case pkg.ErrDelAccountUserNotFound:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInUserNotFound,
			})
		}
	case pkg.ErrDelAccountPasswordIncorrect:
		{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": text.SignInPasswordIncorrect,
			})
		}
	// JWT
	case pkg.ErrRefreshJWTGen:
		{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": text.JWTRefreshErrorGen,
			})
		}
	case pkg.ErrAccessJWTGen:
		{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": text.JWTAccessErrorGen,
			})
		}
	//
	default:
		{
			c.Status(http.StatusInternalServerError)
		}
	}
}
