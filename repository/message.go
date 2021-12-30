package repository

import (
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
	UpdaterID uint       `gorm:"type:int"`
	Updater   User       `gorm:"foreignKey:UpdaterID"`
}
