package routers

import (
	"bin_blog/internal/handler"
)

func RegisterPostRoutes() {
	privPost := privateGroup.Group("/post")
	{
		privPost.GET("/:id", handler.GetUser)
	}

}
