//go:build wireinject

package router

import (
	"webook/internal/handler"
	"webook/internal/repository"
	"webook/internal/repository/cache"
	"webook/internal/repository/database"
	"webook/internal/service"
	"webook/internal/service/oauth2"

	"github.com/google/wire"
)

func InitUserRouter() *UserRouter {
	wire.Build(
		database.Init,
		cache.Init,
		database.NewUserDAO,
		cache.NewUserCache,
		cache.NewCodeCache,
		repository.NewUserRepository,
		service.NewUserService,
		service.NewCodeService,
		handler.NewUserHandler,
		NewUserRouter,
	)
	return new(UserRouter)
}

func InitOAuth2Router() *OAuth2Router {
	wire.Build(
		oauth2.NewOAuth2WechatService,
		handler.NewOAuth2WeChatHandler,
		NewOAuth2Router,
	)
	return new(OAuth2Router)
}
