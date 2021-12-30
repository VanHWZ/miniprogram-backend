package repository

import (
	"time"
)

type Group struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string    `gorm:"default:group"`
	Users       []User    `gorm:"many2many:user_group"`
	Events      []Event   `gorm:"foreignKey:GroupID"`
	Messages    []Message `gorm:"foreignKey:GroupID"`
}

func NextGroup() *Group {
	var group Group
	DB.Create(&group)
	return &group
}
