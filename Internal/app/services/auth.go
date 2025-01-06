package services

import (
	"fmt"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/repositories"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(Username string, Password string) *entities.AuthReturnData {
	var User repositories.User
	if err := User.GetByUsername(Username); err == nil && err != gorm.ErrRecordNotFound {
		return &entities.AuthReturnData{
			Status: http.StatusConflict,
			Message: fmt.Sprintf("ຜູ້ໃຊ້ '%s' ມີໃນລະບົບແລ້ວ", Username),
		}
	}
	if err := User.Create(Username, Password); err != nil {
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
	var User repositories.User
	if err := User.GetByUsername(Username); err != nil {
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

	userPassword := User.Password
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(Password)); err != nil {
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