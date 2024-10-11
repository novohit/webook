package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
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
