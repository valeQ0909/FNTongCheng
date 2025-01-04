package service

import (
	"fntongcheng_back/models"
	"log"
)

type PostServiceImpl struct {
}

// InsertPost 创建一个Post对象
func (psi *PostServiceImpl) InsertPost(post *models.TablePost) bool {
	flag := models.InsertTablePost(post)
	if flag == false {
		log.Println("【Error】 PostServiceImpl 方法InsertPost: 插入失败")
		return false
	}
	return true
}
