package cache

import "github.com/redis/go-redis/v9"

func Init() redis.Cmdable {
	//redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}
