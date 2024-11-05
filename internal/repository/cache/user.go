package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"webook/internal/domain"

	"github.com/redis/go-redis/v9"
)

// A 用到了 B，B一定是接口 => 面向接口
// A 用到了 B，B一定是A的字段 => 规避包变量 包方法（与业务相关的逻辑）
// A 用到了 B，A 绝对不初始化B，而是外面注入
type UserCache struct {
	// client     *redis.Cmdable 这里不能用指针，因为Cmdable是接口，不然获取不到方法
	client     redis.Cmdable
	expiration time.Duration
}

func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		client:     client,
		expiration: time.Minute * 30,
	}
}

func (c *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	val, err := c.client.Get(ctx, c.key(id)).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var user domain.User
	err = json.Unmarshal(val, &user)
	return user, err
}

func (c *UserCache) Set(ctx context.Context, id int64, user domain.User) error {
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, c.key(id), val, c.expiration).Err()
}

func (c *UserCache) key(id int64) string {
	return fmt.Sprintf("user:profile:%d", id)
}

// 建议直接定义成方法而不是函数，因为避免包内其他地方使用
//func key(id int64) string {
//	return fmt.Sprintf("user:profile:%d", id)
//}
