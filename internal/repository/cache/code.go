package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// 编译器会在编译的时候自动将代码注入到变量里面
//
//go:embed lua/set_code.lua
var setCodeScript string

//go:embed lua/verify_code.lua
var verifyCodeScript string

var (
	ErrCodeSendTooMany   = errors.New("验证码发送频繁")
	ErrCodeVerifyTooMany = errors.New("验证次数太多")
	ErrUnknownForCode    = errors.New("发送验证码遇到未知错误")
)

type CodeCache struct {
	client redis.Cmdable
}

func NewCodeCache(client redis.Cmdable) *CodeCache {
	return &CodeCache{client: client}
}

// key = user:biz:code:xxxxx
func (c *CodeCache) SetCode(ctx context.Context, biz string, phone string, code string) error {
	res, err := c.client.Eval(ctx, setCodeScript, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return err
	}

	switch res {
	case 0:
		return nil
	case -1:
		// 发送频繁
		return ErrCodeSendTooMany
	default:
		// res = -2 缓存异常情况
		return ErrUnknownForCode
	}
}

func (c *CodeCache) VerifyCode(ctx context.Context, biz string, phone string, inputCode string) (bool, error) {
	res, err := c.client.Eval(ctx, verifyCodeScript, []string{c.key(biz, phone)}, inputCode).Int()
	if err != nil {
		return false, err
	}

	switch res {
	case 0:
		return true, nil
	case -1:
		return false, ErrCodeVerifyTooMany
	case -2:
		return false, nil
	default:
		return false, ErrUnknownForCode
	}
}

func (c *CodeCache) key(biz string, phone string) string {
	return fmt.Sprintf("user:%s:code:%s", biz, phone)
}
