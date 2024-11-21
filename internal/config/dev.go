//go:build !k8s

package config

// 编译时使用 go build -tags=k8s -o webook 来指定编译哪个文件
// 优势：有编译期的检查
// 缺点：不会动态更新配置文件 每次需要重新编译
var AppConf = AppConfig{
	MySQLConfig: &MySQLConfig{
		DNS: "root:root@tcp(localhost:13306)/webook",
	},
	RedisConfig: &RedisConfig{
		Addr: "127.0.0.1:6379",
	},
	AppId:       "AAAA",
	RedirectUrl: "baidu.com",
}
