package router

import (
	"strings"
	"time"
	"webook/internal/handler"
	"webook/internal/repository"
	"webook/internal/repository/database"
	"webook/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db := initDB()
	uh := initUserHandler(db)
	u := NewUserRouter(uh)

	v1 := r.Group("/api/v1")
	{
		u.RegisterUserRoutes(v1)
	}

	return r
}

func initUserHandler(db *gorm.DB) *handler.UserHandler {
	dao := database.NewUserDAO(db)
	repo := repository.NewUserRepository(dao)
	svc := service.NewUserService(repo)
	uh := handler.NewUserHandler(svc)
	return uh
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook"))
	if err != nil {
		panic("failed to connect database")
	}
	err = database.InitTable(db)
	if err != nil {
		panic("failed to init table")
	}
	return db
}
