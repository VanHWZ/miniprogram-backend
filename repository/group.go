package repository

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Used bool `gorm:"type:bool"`
}

func NewGroupRepo() *gorm.DB {
	return DB.Model(&Group{})
}

func NewGroup() Group {
	var group Group
	result := DB.Model(&Group{}).Where("used=?", false).First(&group)
	if result.RowsAffected == 0 {
		group.Used = false
		DB.Create(&group)
	}
	return group
}
