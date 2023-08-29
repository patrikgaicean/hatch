package main

import (
	"log"

	"github.com/patriuk/hatch/internal/registry"
	"github.com/patriuk/hatch/internal/registry/config"
	"github.com/patriuk/hatch/internal/registry/flags"
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

	err := registry.ListenAndServe(*cfg)
	if err != nil {
		log.Fatal(err)
	}
}
