package main

import (
	"fmt"

	"github.com/patriuk/hatch/internal/gateway"
	"github.com/patriuk/hatch/internal/gateway/config"
	"github.com/patriuk/hatch/internal/gateway/flags"
)

func main() {
	f := flags.ParseFlags()

	cfg := config.New(config.Config{
		Env:  f.Env,
		IP:   f.IP,
		Port: f.Port,
		Redis: config.RedisConfig{
			Host:     f.Redis.Host,
			Port:     f.Redis.Port,
			Password: f.Redis.Password,
		},
	})

	err := gateway.ListenAndServe(*cfg)
	if err != nil {
		fmt.Println("server error")
	}
}
