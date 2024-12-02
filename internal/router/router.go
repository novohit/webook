package router

import (
	"strings"
	"time"
	"webook/internal/config"
	"webook/internal/handler"
	"webook/internal/repository"
	"webook/internal/repository/cache"
	"webook/internal/repository/database"
	"webook/internal/service"
	"webook/pkg/ginx/ratelimit"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(logger.Ginzap(zap.L(), time.DateTime, false))
	//r.Use(logger.RecoveryWithZap(zap.L(), true))
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
	//store := cookie.NewStore([]byte("secret"))
	store, err := sessionRedis.NewStore(16, "tcp", config.AppConf.RedisConfig.Addr, "",
		[]byte("b5ntFvvEfUbKG4Bn"))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("session_id", store))

	r.Use(ratelimit.NewBuilder(initCache(), 5*time.Second, 2).Build())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	//db := initDB()
	//client := initCache()
	//uh := initUserHandler(db, client)
	//u := InitUserRouter()

	v1 := r.Group("/api/v1")
	{
		InitUserRouter().RegisterRoutes(v1)
		InitOAuth2Router().RegisterRoutes(v1)
	}

	return r
}

func initUserHandler(db *gorm.DB, client redis.Cmdable) *handler.UserHandler {
	dao := database.NewUserDAO(db)
	ucache := cache.NewUserCache(client)
	codeCache := cache.NewCodeCache(client)
	repo := repository.NewUserRepository(dao, ucache)
	svc := service.NewUserService(repo)
	codeSvc := service.NewCodeService(codeCache)
	uh := handler.NewUserHandler(svc, codeSvc)
	return uh
}

func initCache() *redis.Client {
	//redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.AppConf.MySQLConfig.DNS))
	if err != nil {
		panic("failed to connect database")
	}
	err = database.InitTable(db)
	if err != nil {
		panic("failed to init table")
	}
	return db
}
