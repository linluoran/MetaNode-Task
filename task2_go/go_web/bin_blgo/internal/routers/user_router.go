package routers

import (
	"bin_blog/internal/handler"
)

func RegisterUserRoutes() {
	pubUser := publicGroup.Group("/user")
	{
		pubUser.POST("/register", handler.registerHandler)
		pubUser.POST("/login", handler.loginHandler)
	}

	privUser := privateGroup.Group("/user")
	{
		privUser.GET("/:id", handler.GetUser)
	}

}
