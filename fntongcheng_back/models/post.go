package models

import (
	"log"
	"time"
)

// 映射字段名
type TablePost struct {
	Id            int64
	Title         string
	PostDetail    string
	Location      string
	Area          int64
	Level         string
	Rent          string
	PayWays       string
	Period        string
	KuanDai       int64
	DianShi       int64
	Shui          int64
	Dian          int64
	DianHua       int64
	KongTiao      int64
	RanQi         int64
	Contact       string
	Tel           string
	HouseCategory string
	PostCategory  string
	PublishDate   time.Time
	ModifyDate    time.Time
}

// 表名映射
func (u TablePost) TableName() string {
	return "posts"
}

// InsertTablePost 将tablePost插入表内
func InsertTablePost(tablePost *TablePost) bool {
	if err := DB.Create(&tablePost).Error; err != nil {
		log.Println("InsertTablePost插入用户失败", err.Error())
		return false
	}
	return true
}
