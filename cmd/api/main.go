package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/patriuk/hatch/internal/api"
	"github.com/patriuk/hatch/internal/api/config"
	"github.com/patriuk/hatch/internal/api/registry"
)

func main() {
	name := flag.String("name", "hatch-service", "Service Name")
	desc := flag.String("desc", "A sample service for testing the hatch gateway", "Service Description")
	env := flag.String("env", "development", "Environment (development|staging|production)")
	registryUrl := flag.String("registry", "", "Registry Service Url")
	flag.Parse()

	protocol := "tcp"
	listener, err := net.Listen(protocol, ":0")
	if err != nil {
		log.Fatal(err)
	}

	// perform heartbeat
	go func() {
		for {
			time.Sleep(5 * time.Second)
			// registry.SendHeartbeat()
		}
	}()

	// register to service registry
	// registry.Register()

	// init config with listener
	cfg, err := config.New(config.Params{
		Name:        *name,
		Description: *desc,
		Env:         *env,
		Protocol:    protocol,
		Listener:    listener,
		RegistryURL: *registryUrl,
	})
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(api.Params{
		Config:   *cfg,
		Listener: listener,
	})

	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}

	// on server <graceful> shutdown, de-register from service registry
	// registry.Unregister()

}

// todo: figure out if this would be better than time.Sleep (seems so..)
func startHeartbeatScheduler() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		registry.SendHeartbeat()
	}
}
