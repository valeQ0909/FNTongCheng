package service

import "fntongcheng_back/models"

// FmtUser 最终封装后,controller返回的FmtUser结构体
type FmtUser struct {
	Id        int64  `json:"id"`
	UserName  string `json:"username"`
	Avatar    string `json:"avatar"`    //头像
	Signature string `json:"signature"` //个性签名
}

type UserService interface {
	// GetUserByUsername 根据username获得User对象
	GetUserByUsername(name string) (models.TableUser, error)
	// GetFmtUserByUsername 根据username获得FmtUser对象
	GetFmtUserByUsername(name string) (FmtUser, error)
	// InsertUser 将user插入表内
	InsertUser(user *models.TableUser) bool
	// GetFmtUserById 未登录情况下,根据user_id获得User对象
	GetFmtUserById(id int64) (FmtUser, error)
	// UpdateAvatar 更新用户头像
	UpdateAvatar(userId int64, avatar string) error
}
