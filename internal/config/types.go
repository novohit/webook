package config

type AppConfig struct {
	*MySQLConfig
	*RedisConfig
}

type MySQLConfig struct {
	DNS string
}

type RedisConfig struct {
	Addr string
}
