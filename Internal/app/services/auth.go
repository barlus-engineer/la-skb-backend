package services

import (
	"fmt"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/repositories"
	"la-skb/Internal/app/text"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth entities.Auth

func (p *Auth) SignUp() *entities.AuthReturnData {
	var User repositories.User
	if err := User.GetByUsername(p.Username); err == nil && err != gorm.ErrRecordNotFound {
		return &entities.AuthReturnData{
			Status:  http.StatusConflict,
			Message: fmt.Sprintf(text.Set["auth.signup.user_exists"], p.Username),
		}
	}
	if err := User.Create(p.Username, p.Password); err != nil {
		return &entities.AuthReturnData{
			Status:  http.StatusInternalServerError,
			Message: text.Set["auth.signup.error"],
		}
	}
	return &entities.AuthReturnData{
		Status:  http.StatusCreated,
		Message: text.Set["auth.signup.success"],
	}
}

func (p *Auth) SignIn() *entities.AuthReturnData {
	var User repositories.User
	if err := User.GetByUsername(p.Username); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.AuthReturnData{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf(text.Set["auth.signin.user_notfound"], p.Username),
			}
		}
		return &entities.AuthReturnData{
			Status:  http.StatusInternalServerError,
			Message: text.Set["auth.signin.server_error"],
		}
	}

	userPassword := User.Password
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(p.Password)); err != nil {
		return &entities.AuthReturnData{
			Status:  http.StatusUnauthorized,
			Message: text.Set["auth.signin.password_incorrect"],
		}
	}

	return &entities.AuthReturnData{
		Status:  http.StatusOK,
		Message: text.Set["auth.signin.success"],
	}
}

func (p *Auth) DeleteAccount() *entities.AuthReturnData {
	var User repositories.User
	if err := User.GetByUsername(p.Username); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.AuthReturnData{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf(text.Set["auth.delete_account.username_notfound"], p.Username),
			}
		}
		return &entities.AuthReturnData{
			Status:  http.StatusInternalServerError,
			Message: text.Set["auth.delete_account.server_error"],
		}
	}

	userPassword := User.Password
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(p.Password)); err != nil {
		return &entities.AuthReturnData{
			Status:  http.StatusUnauthorized,
			Message: text.Set["auth.delete_account.password_incorrect"],
		}
	}

	if err := User.Delete(); err != nil {
		return &entities.AuthReturnData{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf(text.Set["auth.delete_account.error"], p.Username),
		}
	}
	return &entities.AuthReturnData{
		Status:  http.StatusOK,
		Message: fmt.Sprintf(text.Set["auth.delete_account.success"], p.Username),
	}
}
