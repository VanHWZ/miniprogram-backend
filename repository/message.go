package repository

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content string `gorm:"type:text"`
	AuthorRefer uint
	Author User `gorm:"foreignkey:AuthorRefer"`
}
