package service

import (
	"fmt"
	errorCode "miniprogram-backend/errorcode"
	repo "miniprogram-backend/repository"
	repr "miniprogram-backend/representation"
	"miniprogram-backend/util"
	//"strconv"
)

type User struct {
	ID uint `uri:"uid" binding:"required"`
}

type UserRegister struct {
	Name   string `json:"name"    binding:"required"`
	OpenID string `json:"open_id" binding:"required"`
}

type UserUpdate struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required"`
}

type UserAuth struct {
	OpenID string `form:"open_id" binding:"required"`
}

type UserList struct {
	Name     string `form:"name"`
	OpenID   string `form:"open_id"`
	GroupID  uint   `form:"group"`
	Page     int    `form:"page,default=1"`
	PageSize int    `form:"page_size,default=10"`
}

func RegisterUser(userRegister *UserRegister) (*repr.User, *repr.AppError) {
	var cnt int64
	if r := repo.DB.Model(&repo.User{}).Where("open_id=?", userRegister.OpenID).Count(&cnt); r.Error != nil {
		return nil, &repr.AppError{
			Code: errorCode.DatabaseError,
			Message: fmt.Sprintf("Register user(open_id=%s) error: %s", userRegister.OpenID, r.Error.Error()),
		}
	}
	if cnt != 0 {
		return nil, &repr.AppError{
			Code: errorCode.RegisterError,
			Message: fmt.Sprintf("Openid(%s) repeated!", userRegister.OpenID),
		}
	} else {
		nextGroup := repo.NextGroup()
		newUser := repo.User{
			OpenID: userRegister.OpenID,
			Name: userRegister.Name,
			Token: util.NewToken(),
			Groups: []repo.Group{
				*nextGroup,
			},
		}
		if r := repo.DB.Create(&newUser); r.Error != nil {
			return nil, &repr.AppError{
				Code: errorCode.DatabaseError,
				Message: fmt.Sprintf("Register user(open_id=%s) error: %s", userRegister.OpenID, r.Error.Error()),
			}
		}
		return repr.UserSerializer.Serialize(&newUser), nil
	}
}

func UpdateUser(userUpdate *UserUpdate) (*repr.User, *repr.AppError) {
	var user repo.User
	user.ID = userUpdate.ID
	if r := repo.DB.Model(&user).Update("name", userUpdate.Name); r.Error != nil {
		return nil, &repr.AppError{
			Code: errorCode.DatabaseError,
			Message: fmt.Sprintf("User(id=%d) update(name=%s) failed: %s", user.ID, userUpdate.Name, r.Error.Error()),
		}
	}
	if r := repo.DB.First(&user); r.Error != nil {
		return nil, &repr.AppError{
			Code: errorCode.DatabaseError,
			Message: fmt.Sprintf("User(id=%d) update(name=%s) failed: %s", user.ID, userUpdate.Name, r.Error.Error()),
		}
	}
	return repr.UserSerializer.Serialize(&user), nil
}

func ListUser(userInfo *UserList, paginator *util.Paginator) (*[]repr.User, *repr.AppError) {
	var users []repo.User
	db := repo.DB.Model(&repo.User{}).Preload("Groups")
	if userInfo.GroupID != 0 {
		db = db.Joins("JOIN user_group on user_group.user_id = \"user\".id " +
			"JOIN \"group\" on user_group.group_id = \"group\".id " +
			"AND \"group\".id IN (?)", userInfo.GroupID)
	}
	if userInfo.Name != "" {
		db.Where("\"user\".name = ?", userInfo.Name)
	}
	if userInfo.OpenID != "" {
		db = db.Where("\"user\".open_id = ?", userInfo.OpenID)
	}
	if r := db.Group("\"user\".id").
			Scopes(util.Paginate(paginator)).Find(&users); r.Error != nil {
		return nil, &repr.AppError{
			Code: errorCode.DatabaseError,
			Message: fmt.Sprintf("List user error: %s", r.Error.Error()),
		}
	}
	return repr.UserSerializer.SerializeUsers(&users), nil
}

func RetrieveUser(userInfo *User) (*repr.User, *repr.AppError) {
	var user repo.User
	if r := repo.DB.Preload("Groups").First(&user, userInfo.ID); r.Error != nil {
		return nil, &repr.AppError{
			Code: errorCode.DatabaseError,
			Message: fmt.Sprintf("Retrieve user(open_id=%d) error: %s", userInfo.ID, r.Error.Error()),
		}
	}
	return repr.UserSerializer.Serialize(&user), nil
}

func AuthUser(openID string) (*repr.User, *repr.AppError) {
	var user repo.User
	if r := repo.DB.Model(&repo.User{}).Preload("Groups").Where("open_id=?", openID).First(&user); r.Error != nil {
		return nil, &repr.AppError{
			Code: errorCode.DatabaseError,
			Message: fmt.Sprintf("Auth user(open_id=%s) error: %s", openID, r.Error.Error()),
		}
	}
	if user.OpenID == "" {
		return nil, &repr.AppError{
			Code: errorCode.AuthError,
			Message: fmt.Sprintf("Openid(%s) not found.", openID),
		}
	} else {
		tokenString := user.Token
		if !util.ValidateToken(tokenString) {
			tokenString = util.NewToken()
			user.Token = tokenString
			if r := repo.DB.Save(&user); r.Error != nil {
				return nil, &repr.AppError{
					Code: errorCode.DatabaseError,
					Message: fmt.Sprintf("Auth user(open_id=%s) error: %s", openID, r.Error.Error()),
				}
			}
		}
		return repr.UserSerializer.Serialize(&user), nil
	}
}