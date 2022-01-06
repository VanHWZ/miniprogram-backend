package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	errorCode "miniprogram-backend/errorcode"
	rep "miniprogram-backend/representation"
	"miniprogram-backend/service"
	"miniprogram-backend/util"
	"net/http"
	"strconv"
)

func Test(context *gin.Context)  {
	//repo.DB.Create(&repo.GroupInvite{InviterID: 2, GroupID: 2})
}

func ListUser(context *gin.Context) {
	var userRetrieve service.UserList
	if err := context.ShouldBindQuery(&userRetrieve); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	paginator := util.Paginator{Page: userRetrieve.Page, PageSize: userRetrieve.PageSize}
	if users, err := service.ListUser(&userRetrieve, &paginator); err == nil {
		context.JSON(http.StatusOK, &rep.Response{
			Data: gin.H{
				"page": paginator.Page,
				"page_size": paginator.PageSize,
				"data": *users,
			},
		})
	} else {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}
}

func CreateUser(context *gin.Context) {
	var userRegister service.UserRegister
	if err := context.ShouldBind(&userRegister); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if user, err := service.RegisterUser(&userRegister); err == nil {
		context.JSON(http.StatusCreated, &rep.Response{
			Data: *user,
		})
	} else {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}
}

func RetrieveUser(context *gin.Context) {
	var user service.User
	var userSerialized *rep.User
	var err *rep.AppError
	if context.ShouldBindUri(&user) == nil {
		userSerialized, err = service.RetrieveUser(&user)
		if err == nil {
			context.JSON(http.StatusOK, &rep.Response{
				Data: *userSerialized,
			})
			return
		}
	} else {
		*err = rep.AppError{Code: errorCode.ParamError, Message: fmt.Sprintf("Empty openid!")}
	}
	context.JSON(http.StatusBadRequest, err.ErrorResponse())
}

func UpdateUser(context *gin.Context) {
	var userUpdate service.UserUpdate
	if err := context.ShouldBind(&userUpdate); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if userID, err := strconv.Atoi(context.Param("uid")); err == nil {
		userUpdate.ID = uint(userID)
		if user, err := service.UpdateUser(&userUpdate); err == nil {
			context.JSON(http.StatusOK, &rep.Response{
				Data: *user,
			})
		} else {
			context.JSON(http.StatusBadRequest, err.ErrorResponse())
		}
	} else {
		appErr := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appErr.ErrorResponse())
	}

}

func AuthUser(context *gin.Context) {
	var userAuth service.UserAuth
	if context.ShouldBindQuery(&userAuth) == nil {
		user, err := service.AuthUser(userAuth.OpenID)
		if err == nil {
			context.JSON(http.StatusCreated, &rep.Response{
				Data: *user,
			})
		} else {
			context.JSON(http.StatusBadRequest, err.ErrorResponse())
		}
	} else {
		err := rep.AppError{Code: errorCode.ParamError, Message: fmt.Sprintf("Empty openid!")}
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}
}