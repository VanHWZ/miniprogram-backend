package api

import (
	"github.com/gin-gonic/gin"
	errorCode "miniprogram-backend/errorcode"
	rep "miniprogram-backend/representation"
	"miniprogram-backend/service"
	"net/http"
)

func ListMessage(context *gin.Context) {
	var rl service.GroupUserRl
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if messages, err := service.ListMessage(rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
	} else {
		context.JSON(http.StatusOK, &rep.Response{
			Data: gin.H{
				"data": *messages,
			},
		})
	}
}

func CreateMessage(context *gin.Context) {
	var ctn service.MessageCtn
	var rl service.GroupUserRl
	if err := context.ShouldBind(&ctn); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if message, err := service.CreateMessage(&ctn, &rl); err == nil {
		context.JSON(http.StatusCreated, &rep.Response{
			Data: *message,
		})
	} else {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}
}

func RetrieveMessage(context *gin.Context) {
	var rl service.MessageRl
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if message, err := service.RetrieveMessage(&rl); err != nil {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	} else {
		context.JSON(http.StatusOK, &rep.Response{
			Data: *message,
		})
	}
}

func UpdateMessage(context *gin.Context) {
	var messageCtn service.MessageCtn
	var rl service.MessageRl
	if err := context.ShouldBind(&messageCtn); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if message, err := service.UpdateMessage(&messageCtn, &rl); err != nil {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	} else {
		context.JSON(http.StatusOK, &rep.Response{
			Data: *message,
		})
	}
}

func DeleteMessage(context *gin.Context) {
	var rl service.MessageRl
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if err := service.DeleteMessage(&rl); err != nil {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	} else {
		context.Status(http.StatusNoContent)
	}
}