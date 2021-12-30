package middleware

import (
	"github.com/gin-gonic/gin"
	errorCode "miniprogram-backend/errorcode"
	repr "miniprogram-backend/representation"
	"miniprogram-backend/util"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" || !util.ValidateToken(token) {
			// unauthorized
			appError := repr.AppError{Code: errorCode.UnAuthorized, Message: "Invalid or expired openid!"}
			context.JSON(http.StatusUnauthorized, appError.ErrorResponse())
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		context.Next()
	}
}
