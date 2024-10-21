package middleware

import (
	"net/http"
	"time"

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

func (b *AuthMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, ignorePath := range b.ignorePaths {
			if ctx.Request.URL.Path == ignorePath {
				ctx.Next()
				return
			}
		}
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
