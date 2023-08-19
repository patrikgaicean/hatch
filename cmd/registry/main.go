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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", f.Port))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New(config.Params{
		Env: f.Env,
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
