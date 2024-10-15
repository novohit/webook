package service

import (
	"context"
	"time"
	"webook/internal/domain"
	"webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 加密
	// 入库
	u.Created = time.Now().UnixMilli()
	u.Updated = time.Now().UnixMilli()
	return svc.repo.Create(ctx, u)
}
