package repositories

import (
	"encoding/base64"
	"errors"
	"la-skb/Internal/app/database"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User entities.RepoUser

func (p *User) Create(Username string, Password string) error {
	db := database.GetDB()
	Username = base64.StdEncoding.EncodeToString([]byte(Username))
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := models.User{
		Username: Username,
		Password: string(hashPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func (p *User) Delete() error {
	db := database.GetDB()
	if p.ID == 0 {
		return errors.New("ID cannot be zero")
	}
	if err := db.Delete(&models.User{}, p.ID).Error; err != nil {
		return err
	}
	return nil
}

func (p *User) GetByID(ID int) error {
	db := database.GetDB()
	var userModel models.User

	if err := db.First(&userModel, ID).Error; err != nil {
		return err
	}

	p.ID, p.Username, p.Password = userModel.ID, userModel.Username, userModel.Password

	return nil
}

func (p *User) GetByUsername(Username string) error {
	db := database.GetDB()
	Username = base64.StdEncoding.EncodeToString([]byte(Username))
	var userModel models.User

	if err := db.Where("username = ?", Username).First(&userModel).Error; err != nil {
		return err
	}

	p.ID, p.Username, p.Password = userModel.ID, userModel.Username, userModel.Password

	return nil
}

func (p *User) ChangeUsername(newUsername string) error {
	db := database.GetDB()
	var userModel models.User
	if p.ID == 0 {
		return errors.New("ID cannot be zero")
	}
	if err := db.Model(&models.User{}).Where(p.ID).Update("username", newUsername).First(&userModel).Error; err != nil {
		return err
	}
	return nil
}

func (p *User) ChangePassword(newPassword string) error {
	db := database.GetDB()
	var userModel models.User
	if strings.TrimSpace(p.Username) == "" {
		return errors.New("username cannot be empty")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err = db.Model(&models.User{}).Where(p.Username).Update("password", hashPassword).First(&userModel).Error; err != nil {
		return err
	}
	return nil
}