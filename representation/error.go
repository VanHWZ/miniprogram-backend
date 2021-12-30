package representation

import (
	"github.com/gin-gonic/gin"
	error2 "miniprogram-backend/errorcode"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) ErrorResponse() *gin.H {
	// TODO: add string code map
	codeString := error2.CodeMap[e.Code]
	return &gin.H{
		"error_code": codeString,
		"message": e.Message,
	}
}
