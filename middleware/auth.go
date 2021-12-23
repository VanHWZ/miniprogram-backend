package middleware

import (
	"github.com/gin-gonic/gin"
	"miniprogram-backend/util"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" || !util.ValidateToken(token) {
			// unauthorized
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"msg": "",
				"data": "",
			})
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		context.Next()
	}
}
