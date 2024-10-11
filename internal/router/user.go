package router

import (
	"webook/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) RegisterUserRoutes(router *gin.RouterGroup) {
	uh := handler.NewUserHandler()
	userGroup := router.Group("/users")
	{
		userGroup.GET("/profile", uh.Profile) // 获取其他用户信息
		userGroup.POST("/signup", uh.SignUp)  // 注册
		userGroup.POST("/signin", uh.SignIn)  // 登录
		userGroup.PUT("/edit", uh.Edit)       // 登录
	}
}
