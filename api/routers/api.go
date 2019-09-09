package routers

import (
	"github.com/gin-gonic/gin"
	v1 "rs/api/v1"
)

func ApiRouterInit(apiRouterGroup *gin.RouterGroup) {
	apiRouterGroup.GET("/list", v1.ApiList)
	apiRouterGroup.POST("/add", v1.ApiStore)
	apiRouterGroup.POST("/del", v1.ApiDestroyFromID)
}
