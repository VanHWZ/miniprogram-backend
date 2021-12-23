package representation

import (
	"github.com/gin-gonic/gin"
	"strconv"
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
	codeString := strconv.Itoa(e.Code)
	return &gin.H{
		"code": codeString,
		"message": e.Message,
	}
}
