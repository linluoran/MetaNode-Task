package routers

import (
	"bin_blog/internal/handler"
)

func RegisterCommentRoutes() {
	privComment := privateGroup.Group("/comment")
	{
		privComment.POST("/register", handler.registerHandler)
		privComment.POST("/login", handler.loginHandler)
	}
}
