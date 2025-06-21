package routers

import "gin_blog/internal/handler"

func RegisterCommentRoutes() {
	privComment := privateGroup.Group("/comment")
	{
		privComment.POST("/create", handler.CommentCreateHandler)
		privComment.POST("/list", handler.CommentListHandler)
	}
}
