package router

import (
	"webook/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	uh *handler.UserHandler
}

func NewUserRouter(uh *handler.UserHandler) *UserRouter {
	return &UserRouter{uh: uh}
}

func (u *UserRouter) RegisterUserRoutes(router *gin.RouterGroup) {

	userGroup := router.Group("/users")
	{
		userGroup.GET("/profile", u.uh.Profile) // 获取其他用户信息
		userGroup.POST("/signup", u.uh.SignUp)  // 注册
		userGroup.POST("/signin", u.uh.SignIn)  // 登录
		userGroup.PUT("/edit", u.uh.Edit)       // 登录
	}
}
