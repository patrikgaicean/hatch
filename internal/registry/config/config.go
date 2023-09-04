package config

type RedisConfig struct {
	Host     string
	Port     uint16
	Password string
}

type Config struct {
	Env     string
	IP      string
	Port    uint16
	Cleanup int64
	Redis   RedisConfig
}

func New(params Config) *Config {
	return &Config{
		Env:     params.Env,
		IP:      params.IP,
		Port:    params.Port,
		Cleanup: params.Cleanup,
		Redis:   params.Redis,
	}
}
