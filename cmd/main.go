package main

import (
	"fmt"
	"webook/internal/config"
	"webook/internal/router"
	"webook/pkg/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	config.InitV3()
	zap.L().Info("hello world")
	val := viper.GetString("mysql.dns")
	fmt.Println("config", val)
	r := router.InitRouter()
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
