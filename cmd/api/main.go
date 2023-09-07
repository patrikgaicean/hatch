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
	desc := flag.String(
		"desc",
		"sample-service-for-gateway-testing",
		"Service Description",
	)
	env := flag.String(
		"env",
		"development",
		"Environment (development|staging|production)",
	)
	registryAddr := flag.String(
		"registry",
		"http://127.0.0.1:8080",
		"Registry Service Address",
	)
	flag.Parse()

	protocol := "tcp"
	listener, err := net.Listen(protocol, "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	// init config with listener
	cfg, err := config.New(config.Params{
		Name:         *name,
		Description:  *desc,
		Env:          *env,
		Protocol:     protocol,
		Listener:     listener,
		RegistryAddr: *registryAddr,
	})
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(api.Params{
		Config:   *cfg,
		Listener: listener,
	})

	// register to service registry
	registry.Register(*cfg)

	// perform heartbeat -- move this to new func perhaps
	ticker := time.NewTicker(time.Duration(5) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				registry.SendHeartbeat(*cfg)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	err = api.Serve()
	if err != nil {
		log.Fatal(err)
	}

	// on server <graceful> shutdown, de-register from service registry
	// and close heartbeat channel.
	close(quit)
	registry.Unregister(*cfg)
}
