package routers

import (
	"github.com/gin-gonic/gin"
	v1 "rs/api/v1"
	"rs/middelwares/jwt"
)

func FileRouterInit(fileRouterGroup *gin.RouterGroup) {
	fileRouterGroup.Use(jwt.JWTAuth())
	{
		fileRouterGroup.POST("/upload", v1.FileUpload)
	}
}
