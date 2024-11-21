package handler

import (
	"webook/internal/service/oauth2"

	"github.com/gin-gonic/gin"
)

type OAuth2WeChatHandler struct {
	svc *oauth2.OAuth2WechatService
}

func NewOAuth2WeChatHandler(svc *oauth2.OAuth2WechatService) *OAuth2WeChatHandler {
	return &OAuth2WeChatHandler{
		svc: svc,
	}
}

func (handler *OAuth2WeChatHandler) AuthURL(ctx *gin.Context) {
	url, _ := handler.svc.AuthURL(ctx)
	ctx.JSON(200, gin.H{"message": "success", "data": url})
}

func (handler *OAuth2WeChatHandler) CallBack(ctx *gin.Context) {

}
