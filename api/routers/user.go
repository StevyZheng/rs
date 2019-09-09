package routers

import (
	"github.com/gin-gonic/gin"
	v1 "rs/api/v1"
	"rs/middelwares/jwt"
)

func UserRouterInit(userRouterGroup *gin.RouterGroup) {
	userRouterGroup.Use(jwt.JWTAuth())
	{
		userRouterGroup.GET("/list", v1.UserList)
		userRouterGroup.GET("/get/:user_name", v1.UserGetFromName)
		userRouterGroup.POST("/add", v1.UserStore)
		userRouterGroup.POST("/del", v1.UserDestroyFromUserName)
		userRouterGroup.POST("/del/:user_name", v1.UserDestroy)
	}
}
