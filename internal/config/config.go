package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// InitV3 加载远程配置中心
func InitV3() {
	err := viper.AddRemoteProvider("etcd3", "http://127.0.0.1:12379", "/webook")
	viper.SetConfigType("yaml")
	if err != nil {
		panic(err)
	}

	err = viper.ReadRemoteConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err: %v\n", err)
		panic(err)
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("viper.OnConfigChange ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
			fmt.Printf("viper config: %v\n", viper.AllSettings())
		}
	})
}

func InitV2() {
	// 根据命令行参数读取不同环境的配置
	// go run /cmd/main.go --config=./config/dev.yaml
	file := pflag.String("config", "./internal/config/dev.yaml", "config file path")
	pflag.Parse()
	viper.SetConfigFile(*file)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err: %v\n", err)
		panic(err)
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("viper.OnConfigChange ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
			fmt.Printf("viper config: %v\n", viper.AllSettings())
		}
	})
}

func Init() {
	//viper.SetConfigFile("./internal/config/dev.yaml")
	viper.SetConfigName("dev")
	// 按顺序读取 直到读取成功
	viper.AddConfigPath("./internal/config") // optionally look for config in the working directory	err := viper.ReadInConfig()
	viper.AddConfigPath("/etc/appname/")     // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")    // call multiple times to add many search paths

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err: %v\n", err)
		panic(err)
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("viper.OnConfigChange ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
			fmt.Printf("viper config: %v\n", viper.AllSettings())
		}
	})
}
