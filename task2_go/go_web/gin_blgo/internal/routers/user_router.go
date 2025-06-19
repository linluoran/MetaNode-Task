package routers

import "gin_blog/internal/handler"

func RegisterUserRoutes() {
	pubUser := publicGroup.Group("/user")
	{
		pubUser.POST("/register", handler.UserRegisterHandler)
		pubUser.POST("/login", handler.UserLoginHandler)
	}
}
