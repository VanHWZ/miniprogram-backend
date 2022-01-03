package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint       `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string     `gorm:"type:text"`

	GroupID   uint       `gorm:"type:int"`
	CreatorID uint       `gorm:"type:int"`
	Creator   User       `gorm:"foreignKey:CreatorID"`
}

func (m *Message) AfterCreate(tx *gorm.DB) (err error) {
	if r := DB.Preload("Creator").Find(m); r.Error != nil {
		return errors.New("error when creating new message")
	}
	return
}

func (m *Message) AfterUpdate(tx *gorm.DB) (err error) {
	if r := DB.Preload("Creator").Find(m); r.Error != nil {
		return errors.New("error when creating new message")
	}
	return
}
