package router

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	cors.Default()
	r.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// dev
			if strings.Contains(origin, "http://localhost") {
				return true
			}
			return strings.HasPrefix(origin, "https://github.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	//r.Use(middleware.AuthRequire())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	u := &UserRouter{}

	v1 := r.Group("/api/v1")
	{
		u.RegisterUserRoutes(v1)
	}

	return r
}
