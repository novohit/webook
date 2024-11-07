package database

import (
	"webook/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.AppConf.MySQLConfig.DNS))
	if err != nil {
		panic("failed to connect database")
	}
	err = InitTable(db)
	if err != nil {
		panic("failed to init table")
	}
	return db
}
