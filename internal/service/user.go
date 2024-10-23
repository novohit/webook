package service

import (
	"context"
	"time"
	"webook/internal/domain"
	"webook/internal/global"
	"webook/internal/repository"
	"webook/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	// 入库
	u.Created = time.Now().UnixMilli()
	u.Updated = time.Now().UnixMilli()
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) SignIn(ctx context.Context, u domain.User) (*domain.User, error) {
	dbUser, err := svc.repo.GetByEmail(ctx, u.Email)
	if err != nil {
		return nil, global.ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password))
	if err != nil {
		return nil, global.ErrUserOrPassword
	}
	return &dbUser, nil
}

func (svc *UserService) SignInJWT(ctx context.Context, u domain.User, userAgent string) (string, error) {
	dbUser, err := svc.repo.GetByEmail(ctx, u.Email)
	if err != nil {
		return "", global.ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password))
	if err != nil {
		return "", global.ErrUserOrPassword
	}
	token, err := jwt.GenToken(dbUser.Email, userAgent)
	if err != nil {
		return "", err
	}
	return token, nil
}
