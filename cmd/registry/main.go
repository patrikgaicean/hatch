package main

import (
	"fmt"
	"log"
	"net"

	"github.com/patriuk/hatch/internal/registry"
	"github.com/patriuk/hatch/internal/registry/config"
	"github.com/patriuk/hatch/internal/registry/flags"
)

func main() {
	f := flags.ParseFlags()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", f.IP, f.Port))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New(config.Params{
		Env: f.Env,
		Redis: config.RedisConfig{
			Host:     f.Redis.Host,
			Port:     f.Redis.Port,
			Password: f.Redis.Password,
		},
	})

	registry := registry.New(registry.Params{
		Config:   *cfg,
		Listener: listener,
	})

	err = registry.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
