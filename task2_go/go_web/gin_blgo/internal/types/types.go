package types

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type (
	// Response 全局标准化响应
	Response struct {
		Code    int    `form:"code"`
		Data    any    `form:"data"`
		Message string `form:"message"`
	}

	CustomClaims struct {
		ID       uint   `form:"id"`
		Username string `form:"username"`
		jwt.RegisteredClaims
	}
)

type (
	ListReq struct {
		PageSize int `form:"page_size" binding:"omitempty,min=2,max=20"`
		PageNum  int `form:"page_num" binding:"omitempty,min=1"`
	}
)
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

type (
	PostCreateReq struct {
		Title   string `form:"title" binding:"required,min=2,max=20"`
		Content string `form:"content" binding:"required,min=2"`
	}

	PostListReq struct {
		ListReq
		Title string `form:"title" binding:"omitempty,min=1,max=20"`
	}
)

func SuccessRes(c *gin.Context, msg string, data any) {
	c.JSON(
		http.StatusOK,
		Response{
			Code:    http.StatusOK,
			Message: msg,
			Data:    data,
		})
}

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
