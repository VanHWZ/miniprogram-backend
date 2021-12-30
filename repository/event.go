package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

const (
	AnniversaryEventType    = 0
	NotAnniversaryEventType = 1
)

type Event struct {
	ID        uint          `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string        `gorm:"type:text"`
	EventType int           `gorm:"type:int"`
	EventTime time.Time     `gorm:"type:time"`

	GroupID   uint          `gorm:"type:int"`
	CreatorID uint          `gorm:"type:int"`
	Creator   User          `gorm:"foreignKey:CreatorID"`
	UpdaterID uint          `gorm:"type:int"`
	Updater   User          `gorm:"foreignKey:UpdaterID"`
}

func (e *Event) AfterCreate(tx *gorm.DB) (err error) {
	if r := DB.Preload("Creator").Preload("Updater").Find(e); r.Error != nil {
		return errors.New("error when creating new event")
	}
	return
}

func (e *Event) AfterUpdate(tx *gorm.DB) (err error) {
	if r := DB.Preload("Creator").Preload("Updater").Find(e); r.Error != nil {
		return errors.New("error when creating new event")
	}
	return
}
