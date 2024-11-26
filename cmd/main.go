package main

import (
	"fmt"
	"webook/internal/config"
	"webook/internal/router"

	"github.com/spf13/viper"
)

func main() {
	config.InitV3()
	val := viper.GetString("mysql.dns")
	fmt.Println("config", val)
	r := router.InitRouter()
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
