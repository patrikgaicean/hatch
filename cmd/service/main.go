package main

import (
	"log"
	"net"
	"time"

	"github.com/patriuk/hatch/internal/api"
	"github.com/patriuk/hatch/internal/api/registry"
)

func main() {
	listener, err := net.Listen("tcp", ":0")
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

	api := api.New(api.Params{
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
