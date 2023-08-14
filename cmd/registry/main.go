package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/patriuk/hatch/internal/registry"
	"github.com/patriuk/hatch/internal/registry/config"
)

func main() {
	port := flag.Int("port", 8080, "Registry port")
	env := flag.String("env", "development", "Environment (development|staging|production)")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New(config.Params{
		Env: *env,
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
