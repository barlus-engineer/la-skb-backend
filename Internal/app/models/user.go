package models

type User struct {
	General struct {
		ID			uint	`gorm:"primaryKey"`
		Username	string	`gorm:"size:64"`
		Password	string	`gorm:"size:64"`
	}
	Admin struct {
		ID			uint	`gorm:"primaryKey"`
		Username	string	`gorm:"size:64"`
		Password	string	`gorm:"size:64"`
	}
}