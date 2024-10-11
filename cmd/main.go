package main

import (
	"fmt"
	"webook/internal/router"
)

func main() {
	r := router.InitRouter()
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
