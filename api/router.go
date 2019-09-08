package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rs/api/v1"
	"rs/middelwares/jwt"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.NoRoute(api.NotFound)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":     "rs",
			"version": "1.0",
		})
	})
	apiV1 := router.Group("/api/v1")
	apiV1.POST("/login", v1.Login)

	apiDoc := apiV1.Group("/doc")
	{
		apiDoc.GET("/list", v1.ApiList)
		apiDoc.POST("/add", v1.ApiStore)
		apiDoc.POST("/del", v1.ApiDestroyFromID)
	}

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
