package database

import (
	"context"
	"database/sql"

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

func (dao *UserDAO) SelectById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	return u, err
}

type User struct {
	Id int64 `gorm:"primarykey"`
	// https://stackoverflow.com/questions/40092155/difference-between-string-and-sql-nullstring
	// 在多种登录方式下 Email 和 Phone 其中一者可能会为 null，推荐使用sql.Null*表示，而不是 *string
	Email     sql.NullString `gorm:"type:varchar(255);unique"`
	Phone     sql.NullString `gorm:"type:varchar(255);unique"`
	Password  string
	CreatedAt int64
	UpdatedAt int64
}
