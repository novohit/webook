package config

type AppConfig struct {
	*MySQLConfig
	*RedisConfig
	AppId       string
	RedirectUrl string
}

type MySQLConfig struct {
	DNS string
}

type RedisConfig struct {
	Addr string
}
