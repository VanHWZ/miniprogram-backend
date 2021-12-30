package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	errorCode "miniprogram-backend/errorcode"
	repo "miniprogram-backend/repository"
	repr "miniprogram-backend/representation"
	"miniprogram-backend/service"
	"net/http"
)

func GroupPermissionCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		var rl service.GroupUserRl
		var user repo.User
		var group repo.Group
		if err := context.ShouldBindUri(&rl); err != nil {
			appError := repr.AppError{Code: errorCode.DatabaseError, Message: err.Error()}
			context.AbortWithStatusJSON(http.StatusBadRequest, appError.ErrorResponse())
		}
		user.ID = rl.UserID
		group.ID = rl.GroupID
		if err := repo.DB.Model(&user).Association("Groups").Find(&group); err != nil {
			appError := repr.AppError{Code: errorCode.DatabaseError, Message: err.Error()}
			context.AbortWithStatusJSON(http.StatusBadRequest, appError.ErrorResponse())
		}
		if group.Name != "" {
			context.Next()
		} else {
			appError := repr.AppError{Code: errorCode.PermissionDenied, Message: fmt.Sprintf("User(id=%d) has no access to Group(id=%d)", rl.UserID, rl.GroupID)}
			context.AbortWithStatusJSON(http.StatusUnauthorized, appError.ErrorResponse())
		}
	}
}
