package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"webook/pkg/jwt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareBuilder struct {
	ignorePaths []string
}

func NewAuthMiddlewareBuilder() *AuthMiddlewareBuilder {
	return &AuthMiddlewareBuilder{}
}

func (b *AuthMiddlewareBuilder) IgnorePath(ignorePath string) *AuthMiddlewareBuilder {
	b.ignorePaths = append(b.ignorePaths, ignorePath)
	return b
}

func (b *AuthMiddlewareBuilder) IgnorePaths(ignorePaths []string) *AuthMiddlewareBuilder {
	b.ignorePaths = ignorePaths
	return b
}

func (b *AuthMiddlewareBuilder) BuildJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, ignorePath := range b.ignorePaths {
			if ctx.Request.URL.Path == ignorePath {
				ctx.Next()
				return
			}
		}
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "用户未登录",
			})
			return
		}

		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "用户未登录",
			})
			return
		}
		claims, err := jwt.VerifyToken(parts[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "用户未登录",
			})
			return
		}

		userAgent := claims.UserAgent
		if ctx.Request.UserAgent() != userAgent {
			// 不正常的敏感操作 监控 发警告
			// 即使是换设备登录 token 应该是空的
			fmt.Printf("异常登录 当前请求 user-agent [%s], token [%s]\n", ctx.Request.UserAgent(), userAgent)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "用户未登录",
			})
			return
		}

		ctx.Set("user_id", claims.Identify)
	}
}

func (b *AuthMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, ignorePath := range b.ignorePaths {
			if ctx.Request.URL.Path == ignorePath {
				ctx.Next()
				return
			}
		}
		// TODO 使用session-redis鉴权的缺点 用户的每次请求都会访问redis 并发高 redis也可能会扛不住
		session := sessions.Default(ctx)
		userId := session.Get("user_id")
		if userId == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "用户未登录",
			})
			return
		}

		// refresh session
		updated := session.Get("updated_at")
		session.Options(sessions.Options{
			MaxAge: 30 * 60,
		})
		session.Set("user_id", userId)

		now := time.Now().UnixMilli()
		if updated == nil {
			session.Set("updated_at", now)
			session.Save()
			return
		}

		updatedts := updated.(int64)
		if now-updatedts > 10*60*1000 {
			session.Set("updated_at", now)
			session.Save()
		}
	}
}
