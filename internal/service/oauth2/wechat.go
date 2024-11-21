package oauth2

import (
	"context"
	"fmt"
	"net/url"
	"webook/internal/config"

	"github.com/lithammer/shortuuid/v4"
)

type OAuth2WechatService struct {
	appId       string
	redirectUrl string
}

func NewOAuth2WechatService() *OAuth2WechatService {
	return &OAuth2WechatService{
		appId:       config.AppConf.AppId,
		redirectUrl: config.AppConf.RedirectUrl,
	}
}

func (service *OAuth2WechatService) AuthURL(ctx context.Context) (string, error) {
	encodeUrl := url.PathEscape(service.redirectUrl)
	state := shortuuid.New()
	// https://open.weixin.qq.com/connect/qrconnect?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect
	return fmt.Sprintf("https://open.weixin.qq.com/connect/qrconnect?appid=%s"+
		"&redirect_uri=%s&response_type=code&scope=SCOPE&state=%s#wechat_redirect",
		service.appId, encodeUrl, state), nil
}

func (service *OAuth2WechatService) CallBack(ctx context.Context) (url string, err error) {
	return "", nil
}
