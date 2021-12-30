package api

import (
	"github.com/gin-gonic/gin"
	errorCode "miniprogram-backend/errorcode"
	rep "miniprogram-backend/representation"
	"miniprogram-backend/service"
	"net/http"
)

func ListEvent(context *gin.Context) {
	var rl service.GroupUserRl
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if events, err := service.ListEvent(rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	} else {
		context.JSON(http.StatusOK, &rep.Response{
			Data: gin.H{
				"data": *events,
			},
		})
	}
}

func CreateEvent(context *gin.Context) {
	var eventCreate service.EventCtn
	var rl service.GroupUserRl
	if err := context.ShouldBind(&eventCreate); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if event, err := service.CreateEvent(&eventCreate, &rl); err == nil {
		context.JSON(http.StatusCreated, &rep.Response{
			Data: *event,
		})
	} else {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	}
}

func RetrieveEvent(context *gin.Context) {
	var rl service.EventRl
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if event, err := service.RetrieveEvent(&rl); err != nil {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	} else {
		context.JSON(http.StatusOK, &rep.Response{
			Data: *event,
		})
	}
}

func UpdateEvent(context *gin.Context) {
	var eventCtn service.EventCtn
	var rl service.EventRl
	if err := context.ShouldBind(&eventCtn); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if event, err := service.UpdateEvent(&eventCtn, &rl); err != nil {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	} else {
		context.JSON(http.StatusOK, &rep.Response{
			Data: *event,
		})
	}
}

func DeleteEvent(context *gin.Context) {
	var rl service.EventRl
	if err := context.ShouldBindUri(&rl); err != nil {
		appError := rep.AppError{Code: errorCode.ParamError, Message: err.Error()}
		context.JSON(http.StatusBadRequest, appError.ErrorResponse())
		return
	}
	if err := service.DeleteEvent(&rl); err != nil {
		context.JSON(http.StatusBadRequest, err.ErrorResponse())
	} else {
		context.Status(http.StatusNoContent)
	}
}