package services

import (
	"fmt"
	"la-skb/Internal/app/database"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(Username string, Password string) *entities.AuthReturnData {
	db := database.GetDB()

	var user models.User
	if err := db.Where("username = ?", Username).First(&user).Error; err == nil && err != gorm.ErrRecordNotFound {
		return &entities.AuthReturnData{
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
		return &entities.AuthReturnData{
			Status: http.StatusInternalServerError,
			Message: "ມີບັນຫາໃນການສະມັກ",
		}
	}
	return &entities.AuthReturnData{
		Status: http.StatusCreated,
		Message: "ສະມັກສຳເລັດ",
	}
}

func SignIn(Username string, Password string) *entities.AuthReturnData {
	db := database.GetDB()
	var user models.User
	
	err := db.Where("username = ?", Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.AuthReturnData{
				Status: http.StatusNotFound,
				Message: fmt.Sprintf("ບໍ່ມີຜູ້ໃຊ້ '%s' ໃນລະບົບ", Username),
			}
		}
		return &entities.AuthReturnData{
			Status: http.StatusInternalServerError,
			Message: "ເຊີບເວີມີບັນຫາບາງຢ່າງ",
		}
	}

	userPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(Password))
	if err != nil {
		return &entities.AuthReturnData{
			Status: http.StatusUnauthorized,
			Message: "ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ",
		}
	}

	return &entities.AuthReturnData{
		Status: http.StatusOK,
		Message: "ເຂົ້າສູ່ລະບົບສຳເລັດ",
	}
}