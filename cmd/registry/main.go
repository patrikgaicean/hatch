package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/patriuk/hatch/internal/registry"
	"github.com/patriuk/hatch/internal/registry/config"
)

type flags struct {
	port  uint16
	env   string
	redis struct {
		host     string
		port     uint16
		password string
	}
}

func main() {
	var (
		f          flags
		serverPort uint
		redisPort  uint
	)

	flag.UintVar(&serverPort, "port", 8080, "Registry port")
	flag.StringVar(
		&f.env,
		"env",
		"development",
		"Environment (development|staging|production)",
	)
	flag.StringVar(&f.redis.host, "redisHost", "", "Redis Host")
	flag.UintVar(&redisPort, "redisPort", 6379, "Redis Port")
	flag.StringVar(
		&f.redis.password,
		"redisPassword",
		"eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		"Redis Password",
	)
	flag.Parse()

	if serverPort > math.MaxUint16 {
		fmt.Println("Error: Server port value exceeds the range of uint16")
		return
	}
	f.port = uint16(serverPort)

	if redisPort > math.MaxUint16 {
		fmt.Println("Error: Redis port value exceeds the range of uint16")
		return
	}
	f.redis.port = uint16(redisPort)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", f.port))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New(config.Params{
		Env: f.env,
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
