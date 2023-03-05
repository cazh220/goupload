package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goupload/service"
	"goupload/utils"
	"net/http"
	"strconv"
)

const MAX_UPLOAD_SIZE = 1000 * 1000 * 10

// 上传
func FileUpload(context *gin.Context)  {
	file, _ := context.FormFile("file")
	tp := context.PostForm("tp")
	_tp, _ := strconv.Atoi(tp)
	prj := context.PostForm("prj")
	fmt.Println(file.Header)
	// 上传文件限制
	if !utils.EnableUploadFileSize(file.Size) {
		context.JSON(http.StatusOK, gin.H{
			"code": -10001,
			"msg": "上传的文件已超过最大限制",
			"data": nil,
		})
		return
	}

	// 判断文件类型
	if !utils.EnableUploadFileType(file.Header.Get("Content-Type")) {
		context.JSON(http.StatusOK, gin.H{
			"code": -10001,
			"msg": "上传文件类型错误",
			"data": nil,
		})
		return
	}

	fileName := service.DestPath(_tp, prj, file.Filename)

	err := context.SaveUploadedFile(file, fileName)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -10001,
			"msg": "上传失败",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "上传成功",
		"data": fileName,
	})
}