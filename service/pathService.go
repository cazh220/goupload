package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"goupload/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type List struct {
	Path string
	CreateTime string
}

// 获取上传的目录
func DestPath(tp int, prj string, fileName string) string {
	var dir = utils.Conf.FileDir
	switch tp {
	case 1:
		dir += "/videos"
		break
	case 2:
		dir += "/pictures"
		break
	default:
		dir += "/others"
	}

	var projects = []string{"hw", "fl", "yf", "hs", "cc", "yq"}

	if !utils.InArray(prj, projects) {
		dir += "/others"
	} else {
		dir += fmt.Sprintf("%s%s", "/", prj)
	}

	path, _ := utils.PathExists(dir)
	if !path {
		os.MkdirAll(dir, 777)
	}

	return dir + "/" + GenerateName(fileName, prj)
}

// 重新生成名字
func GenerateName(fileName string, prj string) string {
	suffix := filepath.Ext(fileName)
	// 获取当前日期毫秒
	name := utils.Md5Crypt(strconv.FormatInt(time.Now().UnixNano(), 10), prj, fileName)
	name += suffix
	return name
}

// 插入数据
func SaveOne(detectionColl *mongo.Collection, detection interface{}) *mongo.InsertOneResult {
	objId, err := detectionColl.InsertOne(context.TODO(), detection)
	if err != nil {
		log.Fatal(err)
	}

	return objId
}
