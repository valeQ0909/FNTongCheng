package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fntongcheng_backend_go/config"
	"fntongcheng_backend_go/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"
)

type UserServiceImpl struct {
}

// GetUserByUsername 通过name获取User对象
func (usi *UserServiceImpl) GetUserByUsername(name string) (models.TableUser, error) {
	tableUser, err := models.GetTableUserByUsername(name)
	if err != nil {
		log.Println("【Error】 UserServiceImpl:", err.Error())
	}
	return tableUser, nil
}

// GetFmtUserByUsername 通过name获取FmtUser对象
func (usi *UserServiceImpl) GetFmtUserByUsername(name string) (FmtUser, error) {
	fmtUser := FmtUser{
		Id:        0,
		UserName:  "",
		Avatar:    config.DefaultAvatar,
		Signature: config.DefaultSign,
	}
	user, err := models.GetTableUserByUsername(name)

	if err != nil {
		log.Println("【Error】 UserServiceImpl 方法GetFmtUserByUsername:", err.Error())
		return fmtUser, err
	}
	fmtUser.Id = user.Id
	fmtUser.UserName = user.Username
	fmtUser.Avatar = user.Avatar
	fmtUser.Signature = user.Signature

	return fmtUser, nil
}

// InsertUser 创建一个User对象
func (usi *UserServiceImpl) InsertUser(user *models.TableUser) bool {
	flag := models.InsertTableUser(user)
	if flag == false {
		log.Println("【Error】 UserServiceImpl 方法InsertUser: 插入失败")
		return false
	}
	return true
}

// GetFmtUserById 通过Id获取格式化的FmtUser对象
func (usi *UserServiceImpl) GetFmtUserById(id int64) (FmtUser, error) {
	fmtUser := FmtUser{
		Id:        0,
		UserName:  "",
		Avatar:    config.DefaultAvatar,
		Signature: config.DefaultSign,
	}

	user, err := models.GetTableUserById(id)
	if err != nil {
		log.Println("UserServiceImpl 方法GetFmtUserById:", err.Error())
		return fmtUser, err
	}
	fmtUser.Id = user.Id
	fmtUser.UserName = user.Username
	fmtUser.Avatar = user.Avatar
	fmtUser.Signature = user.Signature

	return fmtUser, nil
}

// UpdateAvatar 更新用户头像
func (usi *UserServiceImpl) UpdateAvatar(userId int64, avatar string) error {
	err := models.UpdateAvatar(userId, avatar)
	return err
}

// GenerateToken 根据username生成一个token
func GenerateToken(username string) string {
	u, _ := UserService.GetUserByUsername(new(UserServiceImpl), username)
	token := NewToken(u)
	log.Println("token: ", token)
	return token
}

// NewToken 根据用户信息创建token
func NewToken(user models.TableUser) string {
	expiresTime := time.Now().Unix() + 60*60*24
	id64 := user.Id
	claims := jwt.StandardClaims{
		Audience:  user.Username,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		return token
	} else {
		return "fail"
	}
}

// EnCoder 密码加密
func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
