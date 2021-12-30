package repository

import (
	"time"
)

type User struct {
	ID         uint      `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string    `gorm:"type:varchar(100)"`
	OpenID     string    `gorm:"type:varchar(50)"`
	Token      string    `gorm:"type:varchar(120)"`
	Groups     []Group   `gorm:"many2many:user_group"`
}
