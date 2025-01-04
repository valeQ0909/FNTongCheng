package controller

import (
	"fntongcheng_back/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// SendPost /socialnetwork/user/register/ POST
// 用户注册
func SendPost(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	// 获取帖子相关信息
	title := c.PostForm("title")
	post_detail := c.PostForm("post_detail")
	location := c.PostForm("title")
	area := c.PostForm("title")
	level := c.PostForm("title")
	rent := c.PostForm("title")
	pay_ways := c.PostForm("title")
	period := c.PostForm("title")
	kuan_dai := c.PostForm("title")
	dian_shi := c.PostForm("title")
	shui := c.PostForm("title")
	dian := c.PostForm("title")
	dian_hua := c.PostForm("title")
	kong_tiao := c.PostForm("title")
	ran_qi := c.PostForm("title")
	contact := c.PostForm("title")
	tel := c.PostForm("title")
	house_category := c.PostForm("title")
	post_category := c.PostForm("title")

	// 保存帖子中的图片
	file, err := c.FormFile("image_file")
	//log.Println("file: ", file)
	if err != nil {
		log.Println("用户上传图片失败")
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户上传图片失败",
		})
	}
	user, _ := models.GetTableUserById(userId)
	imagePath := "./media/images/" + user.Username + ".jpg"
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		log.Println("保存图片时发生了错误：", err)
		return
	}

}