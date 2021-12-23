package service

import (
	"fmt"
	repo "miniprogram-backend/repository"
	repr "miniprogram-backend/representation"
	"miniprogram-backend/util"
	"strconv"
)

type UserService struct {}

type UserRegister struct {
	Name string `json:"name" binding:"required"`
	OpenID string `json:"open_id" binding:"required"`
}

type UserAuth struct {
	OpenID string `form:"open_id" binding:"required"`
}

type UserInfo struct {
	Name string `form:"name"`
	OpenID string `form:"open_id"`
	GroupRefer string `form:"group"`
}

func NewUserService() UserService {
	return UserService{}
}

func (service *UserService) Register(userRegister *UserRegister) (*repr.User, *repr.AppError) {
	var cnt int64
	if r := repo.DB.Model(&repo.User{}).Where("open_id=?", userRegister.OpenID).Count(&cnt); r.Error != nil {
		return nil, &repr.AppError{Code: 50, Message: r.Error.Error()}
	}
	if cnt != 0 {
		return nil, &repr.AppError{Code: 1, Message: fmt.Sprintf("Openid(%s) repeated!", userRegister.OpenID)}
	} else {
		newUser := repo.User{
			Name: userRegister.Name,
			OpenID: userRegister.OpenID,
			Token: util.NewToken(),
			GroupRefer: repo.NewGroup().ID,
		}
		if r := repo.DB.Create(&newUser); r.Error != nil {
			return nil, &repr.AppError{Code: 50, Message: r.Error.Error()}
		}
		return repr.UserSerializer.Serialize(&newUser), nil
	}
}

func (service *UserService) Auth(openID string) (*repr.User, *repr.AppError) {
	var user repo.User
	if r := repo.DB.Model(&repo.User{}).Where("open_id=?", openID).First(&user); r.Error != nil {
		return nil, &repr.AppError{Code: 50, Message: r.Error.Error()}
	}
	if user.OpenID == "" {
		return nil, &repr.AppError{Code: 2, Message: fmt.Sprintf("Openid(%s) not found.", openID)}
	} else {
		tokenString := user.Token
		if !util.ValidateToken(tokenString) {
			tokenString = util.NewToken()
			user.Token = tokenString
			if r := repo.DB.Save(&user); r.Error != nil {
				return nil, &repr.AppError{Code: 50, Message: r.Error.Error()}
			}
		}
		return repr.UserSerializer.Serialize(&user), nil
	}
}

func (service *UserService) RetrieveUser(userInfo *UserInfo) (*[]repo.User, *repr.AppError) {
	var users []repo.User
	group, _ := strconv.Atoi(userInfo.GroupRefer)
	userSearchInfo := repo.User{
		Name: userInfo.Name,
		OpenID: userInfo.OpenID,
		GroupRefer: uint(group),
	}
	if r := repo.DB.Model(&repo.User{}).Where(userSearchInfo).Find(&users); r.Error != nil {
		return nil, &repr.AppError{Code: 50, Message: r.Error.Error()}
	}

}