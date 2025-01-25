package services

import (
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/repositories"
	"la-skb/config"
	"la-skb/pkg"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth entities.Auth

/*
SignUp handles user registration.
Requires:
- Auth.Username: The username to be registered.
- Auth.Password: The password for the new user.
*/
func (p *Auth) SignUp() error {
	var User repositories.User
	User.Username, User.Password = p.Username, p.Password
	if err := User.GetByUsername(); err == nil && err != pkg.ErrUserNotFound {
		return pkg.ErrSignupUserExists
	}
	if err := User.Create(); err != nil {
		return pkg.ErrSignup
	}

	return nil
}

/*
SignIn handles user login.
Requires:
- Auth.Username: The username to authenticate.
- Auth.Password: The password for authentication.
*/
func (p *Auth) SignIn() error {
	var User repositories.User
	User.Username, User.Password = p.Username, p.Password
	if err := User.GetByUsername(); err != nil {
		if err == pkg.ErrUserNotFound {
			return pkg.ErrSigninUserNotFound
		}
		return pkg.ErrSignin
	}
	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(p.Password)); err != nil {
		return pkg.ErrSigninPasswordIncorrect
	}

	return nil
}

/*
DeleteAccount handles account deletion.
Requires:
- Auth.Username: The username of the account to delete.
- Auth.Password: The password for authentication.
*/
func (p *Auth) DeleteAccount() error {
	var User repositories.User
	User.Username, User.Password = p.Username, p.Password
	if err := User.GetByUsername(); err != nil {
		if err == pkg.ErrUserNotFound {
			return pkg.ErrDelAccountUserNotFound
		}
		return pkg.ErrDelAccount
	}

	userPassword := User.Password
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(p.Password)); err != nil {
		return pkg.ErrDelAccountPasswordIncorrect
	}

	if err := User.Delete(); err != nil {
		return pkg.ErrDelAccount
	}

	return nil
}

/*
RefreshJWT handles account deletion.
Requires:
- Auth.Username: The username of the account to delete.
*/
func RefreshJWT(c *gin.Context, p Auth) error {
	var cfg = config.LoadConfig()
	claims := jwt.MapClaims{
		"username": p.Username,
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return pkg.ErrRefreshJWTGen
	}

	c.SetCookie(
		"refresh_jwt",
		tokenString,
		math.MaxInt32,
		"/",
		cfg.IP,
		true,
		true,
	)

	return nil
}

/*
AccessJWT handles account deletion.
Requires:
- Auth.Username: The username of the account to delete.
*/
func AccessJWT(p Auth) (string, error) {
	var cfg = config.LoadConfig()
	claims := jwt.MapClaims{
		"username": p.Username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.PublicSecret))
	if err != nil {
		return "", pkg.ErrAccessJWTGen
	}

	return tokenString, nil
}
