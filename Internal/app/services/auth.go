package services

import (
	"fmt"
	"la-skb/Internal/app/database"
	"la-skb/Internal/app/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(Username string, Password string) (bool, string) {
	db := database.GetDB()
	var user models.User
	if err := db.Where("username = ?", Username).First(&user).Error; err != nil {
		return false, fmt.Sprintf("ຜູ້ໃຊ້ '%s' ມີໃນລະບົບແລ້ວ", Username)
	}

	cost := 12
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), cost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	newUser := models.User{
		Username: Username,
		Password: string(hashPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return false, "ມີບັນຫາໃນການສະມັກ"
	}
	return true, "ສະມັກສຳເລັດ"
}

func SignIn(Username string, Password string) (bool, string) {
	db := database.GetDB()
	var user models.User
	
	err := db.Where("username = ?", Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, fmt.Sprintf("ບໍ່ມີຜູ້ໃຊ້ '%s' ໃນລະບົບ", Username)
		}
		return false, "ເຊີບເວີມີບັນຫາບາງຢ່າງ"
	}

	userPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(Password))
	if err != nil {
		return false, "ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ"
	}

	return true, "ເຂົ້າສູ່ລະບົບສຳເລັດ"
}