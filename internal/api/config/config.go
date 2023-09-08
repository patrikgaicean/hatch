package config

import (
	"fmt"
	"log"
	"net"

	"github.com/patriuk/hatch/internal/helpers"
)

// create types for protocol and ipType
type Config struct {
	Env          string
	Name         string
	Address      string
	IP           string
	Port         uint16
	Protocol     string
	IPType       string
	RegistryAddr string
}

type Params struct {
	Env          string
	Name         string
	Protocol     string
	Listener     net.Listener
	RegistryAddr string
}

func New(params Params) (*Config, error) {
	addr := params.Listener.Addr().(*net.TCPAddr)

	ipType, err := helpers.GetIPType(addr.IP.String())
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("GetIPType error: %v", err)
	}

	return &Config{
		Env:          params.Env,
		Name:         params.Name,
		Address:      addr.String(),
		IP:           addr.IP.String(),
		Port:         uint16(addr.Port),
		IPType:       ipType,
		Protocol:     params.Protocol,
		RegistryAddr: params.RegistryAddr,
	}, nil
}
