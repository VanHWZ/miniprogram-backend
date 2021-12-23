package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rep "miniprogram-backend/representation"
	"miniprogram-backend/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{service}
}

func (handler *UserHandler) Retrieve(context *gin.Context) {
	var userInfo service.UserInfo
	if context.BindQuery(&userInfo) == nil {
		fmt.Println(userInfo)
	}
}

func (handler *UserHandler) Create(context *gin.Context)  {
	var userRegister service.UserRegister
	if err := context.Bind(&userRegister); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	user, err := handler.userService.Register(&userRegister)
	if err == nil {
		context.JSON(http.StatusCreated, &rep.Response{
			Data: *user,
		})
	} else {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}

}

func (handler *UserHandler) Auth(context *gin.Context) {
	var userAuth service.UserAuth
	if context.BindQuery(&userAuth) == nil {
		user, err := handler.userService.Auth(userAuth.OpenID)
		if err == nil {
			context.JSON(http.StatusCreated, &rep.Response{
				Data: *user,
			})
		} else {
			context.JSON(http.StatusBadRequest, err.ErrorResponse())
		}
	} else {
		err := rep.AppError{Code: 50, Message: fmt.Sprintf("Empty openid!")}
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}

}