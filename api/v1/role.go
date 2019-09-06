package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rs/models"
)

func RoleList(c *gin.Context) {
	var role models.Role
	result, err := role.RoleList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//添加用户
func RoleStore(c *gin.Context) {
	var role models.Role
	role.RoleName = c.Request.FormValue("rolename")
	id, err := role.RoleInsert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}
