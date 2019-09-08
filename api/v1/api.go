package v1

import (
	"github.com/gin-gonic/gin"
	"rs/models"
	"rs/utils"
)

func ApiList(c *gin.Context) {
	var api models.Api
	result, err := api.ApiList()
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, result, err)
}

//添加角色
func ApiStore(c *gin.Context) {
	var api models.Api
	err := c.ShouldBindJSON(&api)
	//role.RoleName = c.Request.FormValue("url")
	id, err := api.ApiInsert()
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, id, nil)
}

func ApiDestroyFromID(c *gin.Context) {
	var api models.Api
	err := c.ShouldBindJSON(&api.ApiID)
	//roleId, err := strconv.ParseInt(c.Request.FormValue("api_id"), 10, 64)
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	_, err = api.ApiDestroy(api.ApiID)
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, api, nil)
}
