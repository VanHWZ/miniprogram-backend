package router

import (
	"github.com/gin-gonic/gin"
	"miniprogram-backend/handler"
	"miniprogram-backend/service"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	router.POST("/user", userHandler.Create)
	router.GET("/user", userHandler.Retrieve)
	router.GET("/auth", userHandler.Auth)

	return router
}
