package routers

import (
	"gin_blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

var (
	publicGroup  *gin.RouterGroup
	privateGroup *gin.RouterGroup
)

func InitRouter(router *gin.Engine) {
	apiVersion := router.Group("/api/v1")
	{
		// 公开路由组
		publicGroup = apiVersion.Group("")

		// 受保护路由组（需要JWT）
		privateGroup = apiVersion.Group("", middleware.JWTAuthMiddleware())
	}

	RegisterUserRoutes()
	RegisterPostRoutes()
	//RegisterCommentRoutes()
}
