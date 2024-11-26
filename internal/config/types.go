package config

var Conf = new(AppConfig)

type AppConfig struct {
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	AppId        string `mapstructure:"app_id"`
	RedirectUrl  string `mapstructure:"redirect_url"`
}

type MySQLConfig struct {
	DNS string `mapstructure:"dns"`
}

type RedisConfig struct {
	Addr string `mapstructure:"addr"`
}
