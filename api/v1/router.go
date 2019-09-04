package v1

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", Users)
	router.POST("/add_user", Store)
	return router
}
