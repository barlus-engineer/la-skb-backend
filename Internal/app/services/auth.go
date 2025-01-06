package services

import (
	"fmt"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/repositories"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth entities.Auth

func (p *Auth) SignUp() *entities.AuthReturnData {
	var User repositories.User
	if err := User.GetByUsername(p.Username); err == nil && err != gorm.ErrRecordNotFound {
		return &entities.AuthReturnData{
			Status: http.StatusConflict,
			Message: fmt.Sprintf("ມີຊື່ຜູ້ໃຊ້ '%s' ໃນລະບົບແລ້ວ", p.Username),
		}
	}
	if err := User.Create(p.Username, p.Password); err != nil {
		return &entities.AuthReturnData{
			Status: http.StatusInternalServerError,
			Message: "ມີບັນຫາໃນການສະມັກຜູ້້ໃຊ້",
		}
	}
	return &entities.AuthReturnData{
		Status: http.StatusCreated,
		Message: "ສະມັກສຳເລັດແລ້ວ",
	}
}

func (p *Auth) SignIn() *entities.AuthReturnData {
	var User repositories.User
	if err := User.GetByUsername(p.Username); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.AuthReturnData{
				Status: http.StatusNotFound,
				Message: fmt.Sprintf("ບໍ່ມີຊື່ຜູ້ໃຊ້ '%s' ໃນລະບົບ", p.Username),
			}
		}
		return &entities.AuthReturnData{
			Status: http.StatusInternalServerError,
			Message: "ເຊີບເວີມີບັນຫາບາງຢ່າງ",
		}
	}

	userPassword := User.Password
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(p.Password)); err != nil {
		return &entities.AuthReturnData{
			Status: http.StatusUnauthorized,
			Message: "ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ",
		}
	}

	return &entities.AuthReturnData{
		Status: http.StatusOK,
		Message: "ເຂົ້າສູ່ລະບົບສຳເລັດແລ້ວ",
	}
}

func (p *Auth) DeleteAccount() *entities.AuthReturnData {
	var User repositories.User
	if err := User.GetByUsername(p.Username); err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entities.AuthReturnData{
				Status: http.StatusNotFound,
				Message: fmt.Sprintf("ບໍ່ມີຊື່ຜູ້ໃຊ້ '%s' ໃນລະບົບ", p.Username),
			}
		}
		return &entities.AuthReturnData{
			Status: http.StatusInternalServerError,
			Message: "ເຊີບເວີມີບັນຫາບາງຢ່າງ",
		}
	}

	userPassword := User.Password
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(p.Password)); err != nil {
		return &entities.AuthReturnData{
			Status: http.StatusUnauthorized,
			Message: "ລະຫັດຜ່ານບໍ່ຖືກຕ້ອງ",
		}
	}

	if err := User.Delete(); err != nil {
		return &entities.AuthReturnData{
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("ມີບັນຫາໃນການລົບຜູ້ໃຊ້ '%s'", p.Username),
		}
	}
	return &entities.AuthReturnData{
		Status: http.StatusOK,
		Message: fmt.Sprintf("ລົບຜູ້ໃຊ້ '%s' ສຳເລັດແລ້ວ", p.Username),
	}
}