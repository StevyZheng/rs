package v1

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rs/middelwares/jwt"
	"rs/models"
	"rs/utils"
	orm "rs/utils/datebase"
	"strconv"
	"time"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
	models.User
}

func LoginCheck(info LoginInfo) (flag bool, u models.User, err error) {
	var user models.User
	if len(info.Username) == 0 || len(info.Password) == 0 {
		return false, user, nil
	}
	err = orm.Eloquent.Where("user_name = ?", info.Username).Preload("Role").First(&user).Error
	if err != nil {
		return false, user, err
	}
	if info.Password == user.Password {
		return true, user, nil
	} else {
		return false, user, nil
	}
}

func Login(c *gin.Context) {
	var login LoginInfo
	err := c.ShouldBindJSON(&login)
	if err == nil {
		isPass, user, err := LoginCheck(login)
		if isPass {
			generateToken(c, user)
		} else {
			utils.JsonRequest(c, -1, nil, err)
		}
	} else {
		println(err.Error())
		utils.JsonRequest(c, -3, nil, nil)
	}
}

// 生成令牌
func generateToken(c *gin.Context, user models.User) {
	j := &jwt.JWT{
		SigningKey: []byte("newtrekWang"),
	}
	claims := jwt.CustomClaims{
		UserID:   user.UserID,
		UserName: user.UserName,
		Email:    user.Email,
		RoleID:   user.RoleID,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

//列表数据
func UserList(c *gin.Context) {
	var user models.User
	result, err := user.UserList()
	if err != nil {
		utils.JsonRequest(c, -2, nil, err)
		return
	}
	utils.JsonRequest(c, 1, result, nil)
}

type UserStoreInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}

//添加用户
func UserStore(c *gin.Context) {
	var userInfo UserStoreInfo
	err := c.ShouldBindJSON(&userInfo)
	var user models.User
	user.Role.RoleName = userInfo.RoleName
	user.UserName = userInfo.UserName
	user.Password = userInfo.Password
	user.Email = userInfo.Email
	id, err := user.UserInsert()
	if err != nil {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, id, nil)
}

type UserUpdateInfo struct {
	UserID int64 `json:"user_id"`
	UserStoreInfo
}

//修改数据
func UserUpdate(c *gin.Context) {
	var user models.User
	var userUpdateInfo UserUpdateInfo
	err := c.ShouldBind(&userUpdateInfo)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.UserUpdate(id)
	if err != nil || result.UserID == 0 {
		utils.JsonRequest(c, -1, nil, err)
		return
	}
	utils.JsonRequest(c, 1, nil, nil)
}
