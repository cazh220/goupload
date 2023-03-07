package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goupload/service"
	"goupload/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

const MAX_UPLOAD_SIZE = 1000 * 1000 * 10

var dataCollection *mongo.Collection

type Files struct {
	//Id          string `bson:"_id"`
	Path       	string `bson:"path"`
	CreateTime 	string `bson:"create_time"`
}

func init()  {
	clientOptions := options.Client().ApplyURI("mongodb://192.168.144.128:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检测连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Connected to MongoDB!")
	dataCollection = client.Database("resources").Collection("files")
}

// 上传
func FileUpload(context *gin.Context) {
	file, _ := context.FormFile("file")
	tp := context.PostForm("tp")
	_tp, _ := strconv.Atoi(tp)
	prj := context.PostForm("prj")

	// 上传文件限制
	if !utils.EnableUploadFileSize(file.Size) {
		context.JSON(http.StatusOK, gin.H{
			"code": -10001,
			"msg":  "上传的文件已超过最大限制",
			"data": nil,
		})
		return
	}

	// 判断文件类型
	if !utils.EnableUploadFileType(file.Header.Get("Content-Type")) {
		context.JSON(http.StatusOK, gin.H{
			"code": -10001,
			"msg":  "上传文件类型错误",
			"data": nil,
		})
		return
	}

	fileName := service.DestPath(_tp, prj, file.Filename)

	err := context.SaveUploadedFile(file, fileName)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -10001,
			"msg":  "上传失败",
			"data": nil,
		})
		return
	}

	// 写入mongodb
	resObj := service.SaveOne(dataCollection, Files{
		Path: fileName[1:],
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	})

	//fmt.Println("resObj", resObj.InsertedID)
	if resObj.InsertedID == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -10002,
			"msg":  "上传失败",
			"data": nil,
		})
		return
	}


	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": fileName[1:],
	})
}
