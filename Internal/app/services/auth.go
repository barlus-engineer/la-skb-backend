package services

import (
	"la-skb/Internal/app/appError"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/repositories"
	"la-skb/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var cfg = config.LoadConfig()
type Auth entities.Auth

/* 
	SignUp handles user registration.
	Requires:
	- Auth.Username: The username to be registered.
	- Auth.Password: The password for the new user.
*/
func (p *Auth) SignUp() error {
	var User repositories.User
	if err := User.GetByUsername(p.Username); err == nil && err != appError.ErrUserNotFound {
		return appError.ErrSignupUserExists
	}
	if err := User.Create(p.Username, p.Password); err != nil {
		return appError.ErrSignup
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
	if err := User.GetByUsername(p.Username); err != nil {
		if err == appError.ErrUserNotFound {
			return appError.ErrSigninUserNotFound
		}
		return appError.ErrSignin
	}
	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(p.Password)); err != nil {
		return appError.ErrSigninPasswordIncorrect
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
	if err := User.GetByUsername(p.Username); err != nil {
		if err == appError.ErrUserNotFound {
			return appError.ErrDelAccountUserNotFound
		}
		return appError.ErrDelAccount
	}

	userPassword := User.Password
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(p.Password)); err != nil {
		return appError.ErrDelAccountPasswordIncorrect
	}

	if err := User.Delete(); err != nil {
		return appError.ErrDelAccount
	}

	return nil
}

/*
	RefreshJWT handles account deletion.
	Requires:
	- Auth.Username: The username of the account to delete.
*/
func RefreshJWT(p Auth) (string, error) {
	claims := jwt.MapClaims{
		"username": p.Username,
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/*
	AccessJWT handles account deletion.
	Requires:
	- Auth.Username: The username of the account to delete.
*/
func AccessJWT(p Auth) (string, error) {
	claims := jwt.MapClaims{
		"username": p.Username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.PublicSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}