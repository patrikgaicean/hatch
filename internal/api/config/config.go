package config

import (
	"fmt"
	"log"
	"net"

	"github.com/patriuk/hatch/internal/helpers"
)

// create types for protocol and ipType
type Config struct {
	Env         string
	Name        string
	Description string
	Address     string
	Port        int16
	Protocol    string
	IPType      string
	RegistryURL string
}

type Params struct {
	Env         string
	Name        string
	Description string
	Protocol    string
	Listener    net.Listener
	RegistryURL string
}

func New(params Params) (*Config, error) {
	addr := params.Listener.Addr().(*net.TCPAddr)

	ipType, err := helpers.GetIPType(addr.IP.String())
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("GetIPType error: %v", err)
	}

	return &Config{
		Env:         params.Env,
		Name:        params.Name,
		Description: params.Description,
		Address:     addr.String(),
		Port:        int16(addr.Port),
		IPType:      ipType,
		Protocol:    params.Protocol,
		RegistryURL: params.RegistryURL,
	}, nil
}
