package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Host   string
	Port   string
	DbHost string
	DbPort string
	DbDriver string
	DbName string
	FileDir string
	MaxFileSize int64
}
var Conf Config

func init() {
	cfg, err := ini.Load("./config/my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Conf.DbDriver = cfg.Section("db").Key("driver").String()
	Conf.DbHost = cfg.Section("db").Key("host").String()
	Conf.DbName = cfg.Section("db").Key("name").String()
	Conf.DbPort = cfg.Section("db").Key("port").String()
	Conf.Host = cfg.Section("").Key("host").String()
	Conf.Port = cfg.Section("").Key("port").String()
	Conf.FileDir = cfg.Section("file").Key("dir").String()
	arr := strings.Split(cfg.Section("file").Key("max_file_size").String(), " * ")
	a , _ := strconv.Atoi(arr[0])
	b , _ := strconv.Atoi(arr[1])
	c , _ := strconv.Atoi(arr[2])
	Conf.MaxFileSize = int64(a * b * c)
}
