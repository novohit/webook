package repository

import (
	"context"
	"webook/internal/domain"
	"webook/internal/repository/database"
)

type UserRepository struct {
	dao *database.UserDAO
}

func NewUserRepository(dao *database.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
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
		Email:    u.Email,
		Password: u.Password,
		Created:  u.CreatedAt,
		Updated:  u.UpdatedAt,
	}, nil
}
