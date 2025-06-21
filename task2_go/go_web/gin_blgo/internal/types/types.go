package types

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// 通用类型
type (
	// Response 全局标准化响应
	Response struct {
		Code    int    `form:"code"`
		Data    any    `form:"data"`
		Message string `form:"message"`
	}

	// CustomClaims jwt结构体
	CustomClaims struct {
		ID       uint   `form:"id"`
		Username string `form:"username"`
		jwt.RegisteredClaims
	}
)

type (
	// ListReq 全局列表
	ListReq struct {
		PageSize int `form:"page_size" binding:"omitempty,min=2,max=20"`
		PageNum  int `form:"page_num" binding:"omitempty,min=1"`
	}
)

// 用户模块
type (
	UserLoginReq struct {
		Username string `form:"username"  binding:"required,min=3,max=20"`
		Password string `form:"password"  binding:"required,min=8,max=32"`
	}

	UserRegisterReq struct {
		UserLoginReq
		Email string `form:"email" binding:"required,email"`
	}
)

// 文章模块
type (
	PostBaseReq struct {
		ID uint `form:"post_id" binding:"required,min=1"`
	}
	PostCreateReq struct {
		Title   string `form:"title" binding:"required,min=2,max=20"`
		Content string `form:"content" binding:"required,min=2"`
	}

	PostListReq struct {
		ListReq
		Title string `form:"title" binding:"omitempty,min=1,max=20"`
	}

	PostDeleteReq struct {
		PostBaseReq
	}

	PostUpdateReq struct {
		PostBaseReq
		PostCreateReq
	}
)
type (
	CommetCreateReq struct {
		PostBaseReq
		Content string `form:"content" binding:"required,min=1"`
	}

	CommetListReq struct {
		ListReq
		PostID uint `form:"post_id" binding:"required,min=1"`
	}
)

// SuccessRes  全局统一成功响应体
func SuccessRes(c *gin.Context, msg string, data any) {
	c.JSON(
		http.StatusOK,
		Response{
			Code:    http.StatusOK,
			Message: msg,
			Data:    data,
		})
}

// ErrorRes 全局失败响应体
func ErrorRes(c *gin.Context, code int, msg string) {
	statusCode := http.StatusInternalServerError
	if code != 0 {
		statusCode = code
	}
	c.Abort()
	c.JSON(
		http.StatusOK,
		Response{
			Code:    statusCode,
			Message: msg,
		})
}
