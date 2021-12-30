package service

import (
	"fmt"
	"gorm.io/gorm/clause"
	"miniprogram-backend/errorcode"
	repo "miniprogram-backend/repository"
	repr "miniprogram-backend/representation"
	"miniprogram-backend/util"
)

type Group struct {
	ID   string `uri:"gid" binding:"required"`
}

type GroupUserRl struct {
	UserID   uint `uri:"uid" binding:"required"`
	GroupID  uint `uri:"gid" binding:"required"`
}

type GroupList struct {
	Page       int    `form:"page,default=1"`
	PageSize   int    `form:"page_size,default=10"`
}

type GroupInvite struct {
	InviterID   uint `uri:"uid"`
	GroupID     uint `uri:"gid"`
}

func CreateGroup(userID uint) (*repr.Group, *repr.AppError) {
	var group repo.Group
	var user repo.User
	user.ID = userID
	if r := repo.DB.Create(&group); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	if r := repo.DB.First(&user); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	if err := repo.DB.Model(&user).Association("Groups").Append(&group); err != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: err.Error()}
	}
	if r := repo.DB.Preload("Users").Preload("Events").First(&group); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.GroupSerializer.Serialize(&group), nil
}

func RetrieveGroup(groupInfo *Group) (*repr.Group, *repr.AppError) {
	var group repo.Group
	if r := repo.DB.Preload("Users").Preload("Events").Preload("Messages").First(&group, groupInfo.ID); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	return repr.GroupSerializer.Serialize(&group), nil
}

func ListGroup(userID uint, paginator *util.Paginator) (*[]repr.Group, *repr.AppError) {
	var groups []repo.Group
	db := repo.DB.Model(&repo.Group{}).Preload("Users").Preload("Events")
	db = db.Joins("JOIN user_group on user_group.group_id = \"group\".id " +
						"JOIN \"user\" on user_group.user_id = \"user\".id " +
						"AND \"user\".id IN (?)", userID)
	if r := db.Scopes(util.Paginate(paginator)).Find(&groups); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	var groupSerialized []repr.Group
	for _, g := range groups {
		groupSerialized = append(groupSerialized, *repr.GroupSerializer.Serialize(&g))
	}
	return &groupSerialized, nil
}

func EnterGroup(rl *GroupUserRl) (*repr.Group, *repr.AppError) {
	var user repo.User
	var group repo.Group
	user.ID = rl.UserID
	group.ID = rl.GroupID
	if err := repo.DB.Debug().Model(&user).Association("Groups").Append(&group); err != nil{
		return nil, &repr.AppError{Code: errorcode.DatabaseError, Message: err.Error()}
	}
	if r := repo.DB.Preload("Users").Preload("Events").Preload("Messages").Find(&group, rl.GroupID); r.Error != nil {
		return nil, &repr.AppError{Code: errorcode.DatabaseError,Message: r.Error.Error()}
	}
	return repr.GroupSerializer.Serialize(&group), nil
}

func QuitGroup(rl *GroupUserRl) *repr.AppError {
	// TODO: 1.delete group if no users left
	//       2. if
	var user repo.User
	var group repo.Group
	user.ID = rl.UserID
	group.ID = rl.GroupID
	if r := repo.DB.Preload("Groups").Find(&user); r.Error != nil {
		return &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}else if len(user.Groups) == 1 {
		// cannot quit group
		return &repr.AppError{Code: errorcode.GroupQuitError, Message: fmt.Sprintf("Quit error: User(id=%d) belongs to only one Group(id=%d).", rl.UserID, rl.GroupID)}
	}
	if r := repo.DB.Preload("Users").Find(&group); r.Error != nil {
		return &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
	}
	if err := repo.DB.Model(&user).Association("Groups").Delete(group); err != nil{
		return &repr.AppError{Code: errorcode.DatabaseError, Message: err.Error()}
	}
	if len(group.Users) == 1 {
		// delete group and its events and messages
		if r := repo.DB.Select(clause.Associations).Delete(&group); r.Error != nil {
			return &repr.AppError{Code: errorcode.DatabaseError, Message: r.Error.Error()}
		}
	}
	return nil
}
