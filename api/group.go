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

func RetrieveGroup(context *gin.Context) {
	var groupInfo service.Group
	var group *rep.Group
	var appError *rep.AppError
	if err := context.ShouldBindUri(&groupInfo); err == nil {
		if group, appError = service.RetrieveGroup(&groupInfo); appError == nil {
			context.JSON(http.StatusOK, &rep.Response{Data: *group,})
			return
		}
	} else {
		*appError = rep.AppError{
			Code: errorCode.ParamError,
			Message: fmt.Sprintf("Retrieve Group(id=%s) failed: %s",
										context.Param("gid"), err.Error())}
	}
	context.JSON(http.StatusBadRequest, appError.ErrorResponse())
	context.Abort()
}

func ListGroup(context *gin.Context) {
	var groupRetrieve service.GroupList
	var userID uint
	if context.ShouldBindQuery(&groupRetrieve) != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	if uid, err := strconv.Atoi(context.Param("uid")); err != nil {
		appErr := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appErr.ErrorResponse())
	} else {
		userID = uint(uid)
	}
	paginator := util.Paginator{Page: groupRetrieve.Page, PageSize: groupRetrieve.PageSize}
	if groups, err := service.ListGroup(userID, &paginator); err == nil {
		context.JSON(http.StatusOK, &rep.Response{
			Data: gin.H{
				"page": paginator.Page,
				"page_size": paginator.PageSize,
				"data": *groups,
			},
		})
	} else {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}
}

func CreateGroup(context *gin.Context) {
	userID, err := strconv.Atoi(context.Param("uid"))
	if err != nil {
		appErr := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appErr.ErrorResponse())
	}
	if group, err := service.CreateGroup(uint(userID)); err == nil {
		context.JSON(http.StatusCreated, &rep.Response{Data: *group})
	} else {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}
}

func EnterGroup(context *gin.Context) {
	var rl service.GroupUserRl
	var group *rep.Group
	var appError *rep.AppError
	if err := context.ShouldBindUri(&rl); err == nil {
		if group, appError = service.EnterGroup(&rl); appError == nil {
			context.JSON(http.StatusOK, &rep.Response{Data: *group})
			return
		}
	} else {
		*appError = rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
	}
	context.JSON(http.StatusBadRequest, appError.ErrorResponse())
}

func QuitGroup(context *gin.Context) {
	var rl service.GroupUserRl
	var appError *rep.AppError
	if err := context.ShouldBindUri(&rl); err == nil {
		if appError = service.QuitGroup(&rl); appError == nil {
			context.JSON(http.StatusOK, &rep.Response{Data: fmt.Sprintf("Quit group(id=%d) successfully", rl.UserID)})
			return
		}
	} else {
		*appError = rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
	}
	context.JSON(http.StatusBadRequest, appError.ErrorResponse())
}