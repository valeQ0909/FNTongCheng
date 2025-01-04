package service

import "fntongcheng_back/models"

type PostService interface {
	// InsertPost 将post插入表内
	InsertPost(post *models.TablePost) bool
}
