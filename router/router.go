package router

import (
	"github.com/gin-gonic/gin"
	"miniprogram-backend/api"
	"miniprogram-backend/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// user related
	router.GET("/test/", api.Test)
	router.POST("/user/", api.CreateUser)
	router.GET("/user", api.ListUser) // super api
	router.GET("/auth", api.AuthUser)

	jwtAuth := router.Group("/user")
	jwtAuth.Use(middleware.Auth())
	{
		// user related
		jwtAuth.GET("/:uid/", api.RetrieveUser)
		jwtAuth.PUT("/:uid/", api.UpdateUser)

		// group related
		jwtAuth.GET("/:uid/group", api.ListGroup)
		jwtAuth.POST("/:uid/group/", api.CreateGroup)
		jwtAuth.POST("/:uid/group/:gid/enter/", api.EnterGroup)

		groupAuth := jwtAuth.Group("/:uid/group/:gid")
		groupAuth.Use(middleware.GroupPermissionCheck())
		{
			groupAuth.GET("/", api.RetrieveGroup)
			groupAuth.POST("/quit/", api.QuitGroup)

			//message related
			groupAuth.GET("/message", api.ListMessage)
			groupAuth.POST("/message/", api.CreateMessage)
			groupAuth.GET("/message/:mid/", api.RetrieveMessage)
			groupAuth.PUT("/message/:mid/", api.UpdateMessage)
			groupAuth.DELETE("/message/:mid/", api.DeleteMessage)

			//event related
			groupAuth.GET("/event", api.ListEvent)
			groupAuth.POST("/event/", api.CreateEvent)
			groupAuth.GET("/event/:eid/", api.RetrieveEvent)
			groupAuth.PUT("/event/:eid/", api.UpdateEvent)
			groupAuth.DELETE("/event/:eid/", api.DeleteEvent)
		}
	}

	return router
}
