package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(50)"`
	OpenID string `gorm:"type:varchar(50)"`
	Token string `gorm:"type:varchar(120)"`
	GroupRefer uint
	Group Group `gorm:"foreignkey:GroupRefer"`
}

func NewUserRepo() *gorm.DB {
	return DB.Model(&User{})
}