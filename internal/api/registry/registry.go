package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/patriuk/hatch/internal/api/config"
)

// ### service registry data
// 1. Service Name: A unique name or identifier for each service in your registry.
// 2. Service Description: A brief description or metadata about the service.
// 3. Address: The IP address where the service is located (IPv4 or IPv6).
// 4. Port: The port number on which the service is listening.
// 5. Protocol: The network protocol used by the service (e.g., HTTP, TCP, UDP).
// 6. IP Type: The IP address type (IPv4 or IPv6).
// 7. Last Update Time: A timestamp indicating when the service was last updated
// or registered in the registry.
// 8. Additional Metadata: Any additional information you might want to store,
// such as service version, status, or tags

type discovery struct {
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
	IPType   string `json:"ipType"`
}

// Register sends a request to register with the registry service.
func Register(cfg config.Config) {
	payload := &discovery{
		Name:     cfg.Name,
		IP:       cfg.IP,
		Port:     cfg.Port,
		Protocol: cfg.Protocol,
		IPType:   cfg.IPType,
	}
	jsonData, _ := json.Marshal(payload)
	fmt.Println("in register func")
	fmt.Println(cfg.RegistryAddr)

	_, err := http.Post(
		cfg.RegistryAddr+"/register",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("error pls")
		log.Fatal(err)
	}

	// log internally that app registered
}

// Unregister sends a request to unregister from the registry service.
func Unregister(cfg config.Config) {
	payload := &discovery{
		Name:     cfg.Name,
		IP:       cfg.IP,
		Port:     cfg.Port,
		Protocol: cfg.Protocol,
		IPType:   cfg.IPType,
	}

	jsonData, _ := json.Marshal(payload)

	_, err := http.Post(
		cfg.RegistryAddr+"/unregister",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log internally that app unregistered
}

// SendHeartbeat sends a heartbeat request to the registry service.
func SendHeartbeat(cfg config.Config) {
	payload := &discovery{
		Name:     cfg.Name,
		IP:       cfg.IP,
		Port:     cfg.Port,
		Protocol: cfg.Protocol,
		IPType:   cfg.IPType,
	}

	jsonData, _ := json.Marshal(payload)

	_, err := http.Post(
		cfg.RegistryAddr+"/refresh",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
}
