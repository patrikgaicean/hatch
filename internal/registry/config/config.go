package config

type RedisConfig struct {
	Host     string
	Port     uint16
	Password string
}

type Config struct {
	Env   string
	Redis RedisConfig
}

type Params struct {
	Env   string
	Redis RedisConfig
}

func New(params Params) *Config {
	return &Config{
		Env:   params.Env,
		Redis: params.Redis,
	}
}
