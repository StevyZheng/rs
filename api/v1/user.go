package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model "rs/models"
	"strconv"
)

//列表数据
func UserList(c *gin.Context) {
	var user model.User
	user.UserName = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.UserList()

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
func UserStore(c *gin.Context) {
	var user model.User
	user.Role.RoleName = c.Request.FormValue("role")
	user.UserName = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.UserInsert()

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

//修改数据
func UserUpdate(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	user.Password = c.Request.FormValue("password")
	result, err := user.UserUpdate(id)
	if err != nil || result.UserID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}
