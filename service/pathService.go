package service

import (
	"fmt"
	"goupload/utils"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// 获取上传的目录
func DestPath(tp int, prj string, fileName string) string {
	var dir = "./upload"
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
	//name := utils.Md5Crypt(strconv.FormatInt(time.Now().UnixNano(), 10), prj, fileName)
	name := utils.Base64EnCode(strconv.FormatInt(time.Now().UnixNano(), 10))
	name += suffix
	return name
}
