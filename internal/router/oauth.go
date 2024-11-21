package router

import (
	"webook/internal/handler"

	"github.com/gin-gonic/gin"
)

type OAuth2Router struct {
	handler *handler.OAuth2WeChatHandler
}

func NewOAuth2Router(handler *handler.OAuth2WeChatHandler) *OAuth2Router {
	return &OAuth2Router{
		handler: handler,
	}
}

func (r *OAuth2Router) RegisterRoutes(router *gin.RouterGroup) {

	wechat := router.Group("/oauth2/wechat")
	{
		wechat.GET("/authurl", r.handler.AuthURL)
		wechat.Any("/callback", r.handler.CallBack)
	}
}
