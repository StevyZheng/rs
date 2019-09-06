package api

import (
	"github.com/gin-gonic/gin"
	"rs/api/v1"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.NoRoute(api.NotFound)

	apiV1 := router.Group("/api/v1")

	apiV1.GET("/rolelist", v1.RoleList)
	apiV1.POST("roleadd", v1.RoleStore)

	apiV1.GET("/userlist", v1.UserList)
	apiV1.POST("/useradd", v1.UserStore)
	return router
}
