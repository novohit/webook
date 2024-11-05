package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"webook/internal/repository/cache"
)

type CodeService struct {
	cache *cache.CodeCache
}

func NewCodeService(cache *cache.CodeCache) *CodeService {
	return &CodeService{
		cache: cache,
	}
}

func (svc *CodeService) Send(ctx context.Context, biz string, phone string) error {
	// 1. 生成验证码
	code := svc.genCode()
	// 2. 设置redis
	err := svc.cache.SetCode(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	// 3. 调用SMS发送验证码
	// 步骤二成功步骤三失败是否要删除redis的code
	// 不能删除，步骤三有可能是服务商超时，很常见，应该留着缓存限制一分钟后再发送
	return nil
}

func (svc *CodeService) Verify(ctx context.Context, biz string, phone string, inputCode string) (bool, error) {
	ok, err := svc.cache.VerifyCode(ctx, biz, phone, inputCode)
	if errors.Is(err, cache.ErrCodeSendTooMany) {
		// TODO 业务告警
		return false, err
	}
	return ok, err
}

func (svc *CodeService) genCode() string {
	// 1 - 999999
	num := rand.Intn(1000000)
	return fmt.Sprintf("%06d", num)
}
