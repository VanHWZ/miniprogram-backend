package repository

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Content string `gorm:"type:text"`
	AuthorRefer uint
	Author User `gorm:"foreignkey:AuthorRefer"`
	EventType int `gorm:"type:int8"`
}
