package repositories

import (
	"context"
	"fmt"
	"la-skb/Internal/app/cache"
	"la-skb/Internal/app/database"
	"la-skb/Internal/app/entities"
	"la-skb/Internal/app/models"
	"la-skb/config"
	"la-skb/pkg"
	"la-skb/pkg/logger"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User entities.RepoUser

func (p *User) Create() error {
	db := database.GetDB()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := models.User{
		Username: p.Username,
		Password: string(hashPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}
	go cacheUser(newUser)

	return nil
}

func (p *User) Delete() error {
	var (
		db                   = database.GetDB()
		userData models.User = models.User(*p)
	)
	if p.ID == 0 && strings.Trim(p.Username, " ") == "" {
		return pkg.ErrEmptyValue
	}
	if err := db.Delete(&models.User{}, userData.ID).Error; err != nil {
		return err
	}
	if err := cacheDeleteUser(userData); err != nil {
		return err
	}
	return nil
}

func (p *User) GetByID() error {
	var userModel models.User

	if p.ID == 0 {
		return pkg.ErrEmptyValue
	}

	if err := cacheGetByID(p.ID, &userModel); err != nil {
		if err == pkg.ErrUserNotFound {
			return pkg.ErrUserNotFound
		}
		if err := dbGetByID(p.ID, &userModel); err != nil {
			return err
		}
		go cacheUser(userModel)
	}

	p.ID = userModel.ID
	p.Username = userModel.Username
	p.Password = userModel.Password

	return nil
}

func (p *User) GetByUsername() error {
	var (
		userModel models.User
	)

	if p.Username == "" {
		return pkg.ErrEmptyUserName
	}

	if err := cacheGetByUsername(p.Username, &userModel); err != nil {
		if err == pkg.ErrUserNotFound {
			return pkg.ErrUserNotFound
		}
		if err := dbGetByUsername(p.Username, &userModel); err != nil {
			return err
		}
		go cacheUser(userModel)
	}
	p.ID = userModel.ID
	p.Username = userModel.Username
	p.Password = userModel.Password

	return nil
}

// func (p *User) ChangeUsername(newUsername string) error {
// 	db := database.GetDB()
// 	var userModel models.User
// 	if p.ID == 0 {
// 		return error.ErrEmptyValue
// 	}
// 	if err := db.Model(&models.User{}).Where(p.ID).Update("username", newUsername).First(&userModel).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return error.ErrUserNotFound
// 		}
// 		return err
// 	}
// 	return nil
// }

// func (p *User) ChangePassword(newPassword string) error {
// 	db := database.GetDB()
// 	var userModel models.User
// 	if p.ID == 0 {
// 		return error.ErrEmptyValue
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

// =========================================

func dbGetByID(ID uint, mu *models.User) error {
	var (
		db = database.GetDB()
	)
	if err := db.Where("ID = ?", ID).First(&mu).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := cacheIDNotfound(int(ID)); err != nil {
				return err
			}
			return pkg.ErrUserNotFound
		}
		return err
	}
	return nil
}

func dbGetByUsername(Username string, mu *models.User) error {
	var (
		db = database.GetDB()
	)
	if err := db.Where("username = ?", Username).First(&mu).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := cacheUserNotfound(Username); err != nil {
				return err
			}
			return pkg.ErrUserNotFound
		}
		return err
	}
	return nil
}

func cacheUser(p models.User) {
	var (
		ctx     = context.Background()
		rdbPipe = cache.GetCache().Pipeline()
	)

	cacheUserKey := fmt.Sprintf("user:%s", p.Username)
	cacheIDKey := fmt.Sprintf("userid:%d", p.ID)
	cacheValue := pkg.RDBstringify(strconv.Itoa(int(p.ID)), p.Username, p.Password)

	rdbPipe.Set(ctx, cacheUserKey, cacheValue, time.Hour)
	rdbPipe.Set(ctx, cacheIDKey, cacheValue, time.Hour)

	_, err := rdbPipe.Exec(ctx)
	if err != nil {
		logger.Alert("Cannot cache User")
	}
}

func cacheIDNotfound(ID int) error {
	var (
		ctx     = context.Background()
		cfg = config.LoadConfig()
		rdbPipe = cache.GetCache().Pipeline()
	)

	cacheUserKey := fmt.Sprintf("userid:%d", ID)
	cacheValue := "db.null"

	rdbPipe.Set(ctx, cacheUserKey, cacheValue, cfg.CacheTime)

	_, err := rdbPipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func cacheUserNotfound(Username string) error {
	var (
		ctx     = context.Background()
		cfg = config.LoadConfig()
		rdbPipe = cache.GetCache().Pipeline()
	)

	cacheUserKey := fmt.Sprintf("user:%s", Username)
	cacheValue := "db.null"

	rdbPipe.Set(ctx, cacheUserKey, cacheValue, cfg.CacheTime)

	_, err := rdbPipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func cacheDeleteUser(p models.User) error {
	var (
		ctx     = context.Background()
		rdbPipe = cache.GetCache().Pipeline()
	)

	cacheUserKey := fmt.Sprintf("user:%s", p.Username)
	cacheUserIDKey := fmt.Sprintf("userid:%d", p.ID)

	rdbPipe.Del(ctx, cacheUserKey)
	rdbPipe.Del(ctx, cacheUserIDKey)

	_, err := rdbPipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func cacheGetByID(ID uint, p *models.User) error {
	var (
		ctx = context.Background()
		rdb = cache.GetCache()
	)

	cacheIDKey := fmt.Sprintf("userid:%d", ID)
	rdbResult := rdb.Get(ctx, cacheIDKey)
	if rdbResult.Err() != nil {
		return rdbResult.Err()
	}
	if rdbResult.Val() == "db.null" {
		return pkg.ErrUserIDNotFound
	}

	result := pkg.RDBpaser(rdbResult.Val())

	id, err := strconv.Atoi(result[0])
	if err != nil {
		return err
	}

	p.ID = uint(id)
	p.Username = result[1]
	p.Password = result[2]

	return nil
}

func cacheGetByUsername(Username string, p *models.User) error {
	var (
		ctx = context.Background()
		rdb = cache.GetCache()
	)

	cacheUserKey := fmt.Sprintf("user:%s", Username)
	rdbResult := rdb.Get(ctx, cacheUserKey)
	if rdbResult.Err() != nil {
		return rdbResult.Err()
	}
	if rdbResult.Val() == "db.null" {
		return pkg.ErrUserNotFound
	}

	result := pkg.RDBpaser(rdbResult.Val())

	id, err := strconv.Atoi(result[0])
	if err != nil {
		return err
	}
	
	p.ID = uint(id)
	p.Username = result[1]
	p.Password = result[2]

	return nil
}