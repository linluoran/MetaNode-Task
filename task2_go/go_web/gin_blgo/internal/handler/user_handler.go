package handler

import (
	"fmt"
	"gin_blog/internal/config"
	"gin_blog/internal/model"
	"gin_blog/internal/pkg/dao"
	"gin_blog/internal/pkg/logger"
	"gin_blog/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler(c *gin.Context) {
	var user types.UserRegisterReq
	if err := c.ShouldBind(&user); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}

	var userID uint
	result := dao.DB.
		Model(&model.User{}).
		Select("id").
		Where("username = ? or email = ?", user.Username, user.Email).
		Scan(&userID)
	if result.Error != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("用户信息查询失败: %s", result.Error.Error()))
		logger.Log.Error("数据库查询错误: ", zap.Any("err", result.Error))
		return
	}
	if result.RowsAffected > 0 {
		types.ErrorRes(c, 0, "用户名或邮箱已存在")
		return
	}

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("密码哈希报错: %s", err.Error()))
		logger.Log.Error("密码哈希报错: ", zap.Any("err", err))
		return
	}

	err = dao.DB.Create(
		&model.User{
			Username: user.Username,
			Email:    user.Email,
			Password: string(hashedPasswd)}).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("创建用户失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "创建用户成功.", nil)
}

// UserLoginHandler 用户登陆
func UserLoginHandler(c *gin.Context) {
	var user types.UserLoginReq
	if err := c.ShouldBind(&user); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}
	var storedUser model.User
	if err := dao.DB.Where("username = ?", user.Username).
		Take(&storedUser).Error; err != nil {
		types.ErrorRes(c, 0, "用户名或密码错误")
		logger.Log.Error("数据库查询错误: ", zap.Any("err", err))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		types.ErrorRes(c, 0, "用户名或密码错误")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * config.GlobalConfig.Jwt.TokenExpire).Unix(),
		"iss":      config.GlobalConfig.Jwt.Issuer,
	})

	jwtToken, err := token.SignedString([]byte(config.GlobalConfig.Jwt.Secret))
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("Token 生成失败: %s", err.Error()))
		logger.Log.Error("jwt生成错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "用户登录成功", gin.H{"jwt": jwtToken})
}
