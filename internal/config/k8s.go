//go:build k8s

package config

var AppConf = AppConfig{
	&MySQLConfig{
		DNS: "root:root@tcp(localhost:13307)/webook",
	},
	&RedisConfig{
		Addr: "127.0.0.1:16379",
	},
}
