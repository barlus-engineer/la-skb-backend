package services

import (
	"fmt"
	"la-skb/Internal/app/database"
	"la-skb/Internal/app/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ReturnData struct {
	Status	int
	Message	string
}

func SignUp(Username string, Password string) *ReturnData {
	db := database.GetDB()

	var user models.User
	if err := db.Where("username = ?", Username).First(&user).Error; err == nil && err != gorm.ErrRecordNotFound {
		return &ReturnData{
			Status: http.StatusConflict,
			Message: fmt.Sprintf("ຜູ້ໃຊ້ '%s' ມີໃນລະບົບແລ້ວ", Username),
		}
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	newUser := models.User{
		Username: Username,
		Password: string(hashPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return &ReturnData{
			Status: http.StatusInternalServerError,
			Message: "ມີບັນຫາໃນການສະມັກ",
		}
	}
	return &ReturnData{
		Status: http.StatusCreated,
		Message: "ສະມັກສຳເລັດ",
	}
}

func SignIn(Username string, Password string) *ReturnData {
	db := database.GetDB()
	var user models.User
	
	err := db.Where("username = ?", Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &ReturnData{
				Status: http.StatusNotFound,
				Message: fmt.Sprintf("ບໍ່ມີຜູ້ໃຊ້ '%s' ໃນລະບົບ", Username),
			}
		}
		return &ReturnData{
			Status: http.StatusInternalServerError,
			Message: "ເຊີບເວີມີບັນຫາບາງຢ່າງ",
		}
	}

	userPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(Password))
	if err != nil {
		return &ReturnData{
			Status: http.StatusUnauthorized,
			Message: "ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ",
		}
	}

	return &ReturnData{
		Status: http.StatusOK,
		Message: "ເຂົ້າສູ່ລະບົບສຳເລັດ",
	}
}