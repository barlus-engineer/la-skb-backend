package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:64;unique"`
	Password string `gorm:"size:64"`
}