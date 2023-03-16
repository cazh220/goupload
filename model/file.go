package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Files struct {
	Id	string			`bson:"_id"`
	Path string 		`bson:"path"`
	Size int64  		`bson:"size"`
	Tp int				`bson:"tp"`
	Prj string			`bson:"prj"`
	CreateTime string	`bson:"create_time"`
}

// 获取文件列表
func GetFilesList(detectionColl *mongo.Collection, filter bson.D, limit int64, page int64) []*Files  {
	//filter := bson.D{{"tp", 3}}
	opts := options.Find().SetSort(bson.D{{"create_time", -1}})

	offset := (page -1)*limit
	if limit > 0 {
		opts = opts.SetLimit(int64(limit)).SetSkip(int64(offset))
	}

	fmt.Println(opts)
	cur, err := detectionColl.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	var files []*Files
	err = cur.All(context.Background(), &files)
	if err != nil {
		log.Fatal(err)
	}
	_ = cur.Close(context.Background())
	return files
}

// 获取总数
func GetFilesNum(detectionColl *mongo.Collection, filter bson.D) int {
	count, err := detectionColl.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return 0
	}

	return int(count)
}


