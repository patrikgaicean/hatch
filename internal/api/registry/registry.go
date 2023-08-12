package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "github.com/patriuk/hatch/internal/api/config"
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
type register struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	details
}

type details struct {
	Address  string `json:"address"`
	Port     int16  `json:"port"`
	Protocol string `json:"protocol"`
	IPType   string `json:"ipType"`
}

// Register sends a request to register with the registry service.
func Register() {

	// probably need to send this directly from where it's called?
	// i.e. main
	payload := &register{
		Name:        "sample-service",
		Description: "service to test the app interactions",
		details: details{
			Address:  "need a param for this",
			Port:     666, // need param
			Protocol: "need a param for this",
			IPType:   "need a param for this",
		},
	}
	jsonData, _ := json.Marshal(payload)
	fmt.Println(string(jsonData))

	// cfg := config.New()

	_, err := http.Post(
		// cfg.RegistryURL,
		"someurl",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log internally that app registered
}

// Unregister sends a request to unregister from the registry service.
func Unregister() {
	// Your logic to perform the "unregister" HTTP request.
	payload := &details{
		Address:  "hithere",
		Port:     666,
		Protocol: "abc",
		IPType:   "string",
	}

	jsonData, _ := json.Marshal(payload)

	// cfg := config.New()

	_, err := http.Post(
		// cfg.RegistryURL,
		"someurl",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log internally that app unregistered
}

// SendHeartbeat sends a heartbeat request to the registry service.
func SendHeartbeat() {
	// Your logic to perform the "heartbeat" HTTP request.
	payload := details{
		Address:  "need a param for this",
		Port:     666, // need param
		Protocol: "need a param for this",
		IPType:   "need a param for this",
	}
	jsonData, _ := json.Marshal(payload)
	fmt.Println(string(jsonData))

	// cfg := config.New()

	_, err := http.Post(
		// cfg.RegistryURL,
		"someurl",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
}
