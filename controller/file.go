package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goupload/model"
	"goupload/service"
	"goupload/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

var dataCollection *mongo.Collection

type Files struct {
	Path       	string 	`bson:"path"`
	Size		int64	`bson:"size"`
	Tp			int		`bson:"tp"`
	Prj			string	`bson:"prj"`
	CreateTime 	string 	`bson:"create_time"`
}

func init()  {
	conStr := utils.Conf.DbDriver+"://"+utils.Conf.DbHost+":"+utils.Conf.DbPort
	clientOptions := options.Client().ApplyURI(conStr)
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
	dataCollection = client.Database(utils.Conf.DbName).Collection("files")
}

/**
 * 上传文件
 */
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
		Size: file.Size,
		Tp: _tp,
		Prj: prj,
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

/**
 * 查看图片列表
 * id 根据id
 * type  根据type
 * prj  根据项目
 * 分页排序
 */
func ViewFiles(context *gin.Context)  {
	var files model.Files
	if result := context.ShouldBindJSON(&files); result != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -10003,
			"msg":  "参数错误",
			"data": nil,
		})
		return
	}

	//filter := bson.D{{"tp", 2}}
	filter := bson.D{}

	list := model.GetFilesList(dataCollection, filter)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查看成功",
		"data": list,
	})
}
