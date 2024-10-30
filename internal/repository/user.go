package repository

import (
	"context"
	"webook/internal/domain"
	"webook/internal/repository/cache"
	"webook/internal/repository/database"
)

type UserRepository struct {
	dao   *database.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(dao *database.UserDAO, cache *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: cache,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, database.User{
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.Created,
		UpdatedAt: u.Updated,
	})
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.SelectByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Created:  u.CreatedAt,
		Updated:  u.UpdatedAt,
	}, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	// 1. query cache
	user, err := r.cache.Get(ctx, id)
	// TODO 思考 是判断 cacheUser 还是 err
	// 只要err为nil 就认为缓存有数据 即使为空
	if err == nil {
		return user, nil
	}

	// 如果 err 不等于记录不存在 说明 redis 可能崩了
	// 此时是否继续往下走查询DB，如果继续 怎么保护数据库不被大流量（限流）

	dbUser, err := r.dao.SelectById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	user = domain.User{
		Id:       dbUser.Id,
		Email:    dbUser.Email,
		Password: dbUser.Password,
		Created:  dbUser.CreatedAt,
		Updated:  dbUser.UpdatedAt,
	}

	// 由于用缓存本身一般就做不到强一致性 所以下面代码可以放到 goroutine
	// 设置缓存
	err = r.cache.Set(ctx, id, user)
	if err != nil {
		// 缓存设置失败 打日志 做监控
	}
	return user, err
}
