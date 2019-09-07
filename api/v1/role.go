package v1

import (
	"github.com/gin-gonic/gin"
	"rs/models"
	"rs/utils"
)

func RoleList(c *gin.Context) {
	var role models.Role
	result, err := role.RoleList()
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, result, err)
}

//添加用户
func RoleStore(c *gin.Context) {
	var role models.Role
	err := c.ShouldBindJSON(&role.RoleName)
	//role.RoleName = c.Request.FormValue("rolename")
	id, err := role.RoleInsert()

	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, id, nil)
}

func RoleDestroyFromID(c *gin.Context) {
	var role models.Role
	err := c.ShouldBindJSON(&role.RoleID)
	//roleId, err := strconv.ParseInt(c.Request.FormValue("role_id"), 10, 64)
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	_, err = role.RoleDestroy(role.RoleID)
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, role, nil)
}

func RoleDestroyFromRoleName(c *gin.Context) {
	var role models.Role
	err := c.ShouldBindJSON(&role.RoleName)
	//role.RoleName = c.Request.FormValue("rolename")
	roleId := role.GetRoleIDFromRoleName(role.RoleName)
	_, err = role.RoleDestroy(roleId)
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, role, nil)
}
