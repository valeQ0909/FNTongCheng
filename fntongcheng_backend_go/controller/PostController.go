package controller

import (
	"github.com/gin-gonic/gin"
)

type PostResponse struct {
	Response
}
type PostsListResponse struct {
	Response
}

// Feed 获取首页的帖子
func Feed(c *gin.Context) {

}

// SendPost 发送帖子
func SendPost(c *gin.Context) {

}

// PostList 获取帖子列表  /post/postlist/
func PostList(c *gin.Context) {

}

// PostDetail 帖子详情
func PostDetail(c *gin.Context) {

}
