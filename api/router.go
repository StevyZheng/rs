package api

import (
	"github.com/gin-gonic/gin"
	"rs/api/v1"
	"rs/middelwares/jwt"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.NoRoute(api.NotFound)

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/login", v1.Login)

	apiRole := apiV1.Group("/role")
	apiRole.Use(jwt.JWTAuth())
	{
		apiRole.GET("/list", v1.RoleList)
		apiRole.POST("/add", v1.RoleStore)
	}

	apiUser := apiV1.Group("/user")
	apiUser.Use(jwt.JWTAuth())
	{
		apiUser.GET("/list", v1.UserList)
		apiUser.POST("/add", v1.UserStore)
	}
	return router
}
