package repositories

import (
	"context"
	"encoding/base64"
	"fmt"
	"la-skb/Internal/app/appError"
	"la-skb/Internal/app/cache"
	"la-skb/Internal/app/database"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/models"
	"la-skb/pkg"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User entities.RepoUser

func (p *User) Create() error {
	ctx := context.Background()
	db := database.GetDB()
	rdb := cache.GetCache().Pipeline()
	Username := base64.StdEncoding.EncodeToString([]byte(p.Username))
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
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

	cacheUserKey := fmt.Sprintf("user:%s", newUser.Username)
	cacheUserValue := pkg.RDBstringify(strconv.FormatUint(uint64(newUser.ID), 10), newUser.Username, newUser.Password)
	cacheUserIDKey := fmt.Sprintf("user:%d", newUser.ID)
	cacheUserIDValue := pkg.RDBstringify(strconv.FormatUint(uint64(newUser.ID), 10), newUser.Username, newUser.Password)

	rdb.Set(ctx, cacheUserKey, cacheUserValue, time.Hour)
	rdb.Set(ctx, cacheUserIDKey, cacheUserIDValue, time.Hour)

	_, err = rdb.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *User) Delete() error {
	ctx := context.Background()
	db := database.GetDB()
	rdb := cache.GetCache()
	if p.ID == 0 {
		return appError.ErrEmptyValue
	}
	if err := db.Delete(&models.User{}, p.ID).Error; err != nil {
		return err
	}
	cacheUserIDKey := fmt.Sprintf("user:%d", p.ID)
	if err := rdb.Del(ctx, cacheUserIDKey).Err(); err != nil {
		return err
	}
	return nil
}

func (p *User) GetByID() error {
	ctx := context.Background()
	db := database.GetDB()
	rdb := cache.GetCache()
	var userModel models.User
	if p.ID == 0 {
		return appError.ErrEmptyValue
	}
	
	cacheUserIDKey := fmt.Sprintf("user:%d", p.ID)
	rdbResult := rdb.Get(ctx, cacheUserIDKey)
	if rdbResult.Err() == nil {
		fmt.Println(rdbResult.Val())
	}

	if err := db.First(&userModel, p.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return appError.ErrUserIDNotFound
		}
		return err
	}

	decodedUsername, err := base64.StdEncoding.DecodeString(userModel.Username)
	if err != nil {
		return err
	}
	p.ID, p.Username, p.Password = userModel.ID, string(decodedUsername), userModel.Password

	return nil
}

func (p *User) GetByUsername() error {
	ctx := context.Background()
	db := database.GetDB()
	rdb := cache.GetCache()

	if p.Username == "" {
		return appError.ErrEmptyValue
	}
	Username := base64.StdEncoding.EncodeToString([]byte(p.Username))
	var userModel models.User

	cacheUserKey := fmt.Sprintf("user:%s", Username)
	rdbResult := rdb.Get(ctx, cacheUserKey)
	if rdbResult.Err() == nil {
		result := pkg.RDBpaser(rdbResult.Val())
		fmt.Print(rdbResult.Val())

		decodedUsername, err := base64.StdEncoding.DecodeString(result[1])
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(result[0])
		if err != nil {
			return err
		}
		p.ID, p.Username, p.Password = uint(id), string(decodedUsername), result[2]

		return nil
	}

	// If its else

	if err := db.Where("username = ?", Username).First(&userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return appError.ErrUserNotFound
		}
		return err
	}

	rdbPipe := rdb.Pipeline()

	cacheUserKey = fmt.Sprintf("user:%s", userModel.Username)
	cacheUserValue := fmt.Sprintf("%v", userModel)
	cacheUserIDKey := fmt.Sprintf("user:%d", userModel.ID)
	cacheUserIDValue := fmt.Sprintf("%v", userModel)

	rdbPipe.Set(ctx, cacheUserKey, cacheUserValue, time.Hour)
	rdbPipe.Set(ctx, cacheUserIDKey, cacheUserIDValue, time.Hour)

	_, err := rdbPipe.Exec(ctx)
	if err != nil {
		return err
	}

	decodedUsername, err := base64.StdEncoding.DecodeString(userModel.Username)
	if err != nil {
		return err
	}
	p.ID, p.Username, p.Password = userModel.ID, string(decodedUsername), userModel.Password

	return nil
}

// func (p *User) ChangeUsername(newUsername string) error {
// 	db := database.GetDB()
// 	var userModel models.User
// 	if p.ID == 0 {
// 		return appError.ErrEmptyValue
// 	}
// 	if err := db.Model(&models.User{}).Where(p.ID).Update("username", newUsername).First(&userModel).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return appError.ErrUserNotFound
// 		}
// 		return err
// 	}
// 	return nil
// }

// func (p *User) ChangePassword(newPassword string) error {
// 	db := database.GetDB()
// 	var userModel models.User
// 	if p.ID == 0 {
// 		return appError.ErrEmptyValue
// 	}
// 	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	if err = db.Model(&models.User{}).Where(p.ID).Update("password", hashPassword).First(&userModel).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }