package config

type Config struct {
	Env string
}

type Params struct {
	Env string
}

func New(params Params) *Config {
	return &Config{
		Env: params.Env,
	}
}
