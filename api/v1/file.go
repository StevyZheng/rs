package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func GetConf(confName string) {

}

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	//fileName := file.Filename
	savePath := path.Join("C:\\Users\\Stevy\\Desktop\\", file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	c.String(http.StatusOK, "上传文件成功")
}
