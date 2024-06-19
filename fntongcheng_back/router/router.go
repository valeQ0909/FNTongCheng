package router

import (
	"fntongcheng_back/controller"
	"fntongcheng_back/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/fntongchengapiv1/static", "./static")
	r.Static("/fntongchengapiv1/media", "./media")

	r.GET("/test/", controller.Test)

	apiRouter := r.Group("/fntongchengapiv1")
	/*
		用户
	*/
	//用户注册
	apiRouter.POST("/user/register/", controller.Register)
	//用户登录
	apiRouter.POST("/user/login/", controller.Login)
	//用户信息
	apiRouter.POST("/user/info/", utils.Auth(), controller.UserInfo)
	//根据token返回用户信息
	apiRouter.GET("/user/token/", utils.Auth(), controller.Token)
	//
	apiRouter.POST("/user/updateavatar/", utils.Auth(), controller.UpdateAvatar)
	//
	apiRouter.POST("/user/updatepersonalinfo/", utils.Auth(), controller.UpdatePersonalInfo)

}
