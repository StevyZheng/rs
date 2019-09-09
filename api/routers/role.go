package routers

import (
	"github.com/gin-gonic/gin"
	v1 "rs/api/v1"
	"rs/middelwares/jwt"
)

func RoleRouterInit(roleRouterGroup *gin.RouterGroup) {
	roleRouterGroup.Use(jwt.JWTAuth())
	{
		roleRouterGroup.GET("/list", v1.RoleList)
		roleRouterGroup.GET("/get/:role_name", v1.RoleGetFromName)
		roleRouterGroup.POST("/add", v1.RoleStore)
		roleRouterGroup.POST("/del/:role_name", v1.RoleDestroyFromRoleName)
	}
}
