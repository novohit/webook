package database

import (
	"context"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	return dao.db.WithContext(ctx).Create(&u).Error
}

func (dao *UserDAO) SelectByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
}

type User struct {
	Id        int64  `gorm:"primarykey"`
	Email     string `gorm:"type:varchar(255);unique"`
	Password  string
	CreatedAt int64
	UpdatedAt int64
}
