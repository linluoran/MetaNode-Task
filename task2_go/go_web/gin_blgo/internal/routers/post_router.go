package routers

import "gin_blog/internal/handler"

func RegisterPostRoutes() {
	privPost := privateGroup.Group("/post")
	{
		privPost.POST("/create", handler.PostCreateHandler)
		privPost.POST("/list", handler.PostListHandler)
	}

}
