package router

import (
	"webook/internal/handler"
	"webook/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	uh *handler.UserHandler
}

func NewUserRouter(uh *handler.UserHandler) *UserRouter {
	return &UserRouter{uh: uh}
}

func (u *UserRouter) RegisterUserRoutes(router *gin.RouterGroup) {

	requireLogin := middleware.NewAuthMiddlewareBuilder().IgnorePath("/api/v2/users/profile").BuildJWT()
	userGroup := router.Group("/users")
	{
		userGroup.GET("/profile", requireLogin, u.uh.Profile) // 获取其他用户信息
		userGroup.POST("/signup", u.uh.SignUp)                // 注册
		userGroup.POST("/logout", u.uh.SignOut)               // 注销
		userGroup.POST("/login", u.uh.SignIn)                 // 登录
		userGroup.POST("/login-jwt", u.uh.SignInJWT)          // 登录
		userGroup.PUT("/edit", requireLogin, u.uh.Edit)       // 编辑
		userGroup.POST("/login/send-sms", u.uh.SendSMS)       // 发送验证码
		userGroup.POST("/login/sms", u.uh.SignInSMS)          // 验证码登录
	}
}
