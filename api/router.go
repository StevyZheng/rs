package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rs/api/routers"
	"rs/api/v1"
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
	apiV1.POST("/register", v1.UserStore)

	apiDoc := apiV1.Group("/doc")
	routers.ApiRouterInit(apiDoc)

	apiRole := apiV1.Group("/role")
	routers.RoleRouterInit(apiRole)

	apiUser := apiV1.Group("/user")
	routers.UserRouterInit(apiUser)

	return router
}
