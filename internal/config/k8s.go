//go:build k8s

package config

var AppConf = AppConfig{
	MySQLConfig: &MySQLConfig{
		DNS: "root:root@tcp(localhost:13307)/webook",
	},
	RedisConfig: &RedisConfig{
		Addr: "127.0.0.1:16379",
	},
	AppId:       "AAAA",
	RedirectUrl: "baidu.com",
}
