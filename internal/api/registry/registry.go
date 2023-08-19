package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/patriuk/hatch/internal/api/config"
	"github.com/patriuk/hatch/internal/helpers"
)

type discovery struct {
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
	IPType   string `json:"ipType"`
}

func Register(cfg config.Config) {
	payload := &discovery{
		Name:     cfg.Name,
		IP:       cfg.IP,
		Port:     cfg.Port,
		Protocol: cfg.Protocol,
		IPType:   cfg.IPType,
	}

	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest(
		http.MethodPut,
		cfg.RegistryAddr+"/register",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("successfully registered")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	errs := make(map[string]interface{})
	err = json.Unmarshal(body, &errs)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalf(
		"failed to register\ncode %d\nerrors: %v",
		resp.StatusCode,
		helpers.PrettyPrint(errs),
	)
}

func Unregister(cfg config.Config) {
	payload := &discovery{
		Name:     cfg.Name,
		IP:       cfg.IP,
		Port:     cfg.Port,
		Protocol: cfg.Protocol,
		IPType:   cfg.IPType,
	}

	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest(
		http.MethodDelete,
		cfg.RegistryAddr+"/unregister",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("successfully unregistered")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	errs := make(map[string]interface{})
	err = json.Unmarshal(body, &errs)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalf(
		"failed to unregister\ncode %d\nerrors: %v",
		resp.StatusCode,
		helpers.PrettyPrint(errs),
	)
}

func SendHeartbeat(cfg config.Config) {
	payload := &discovery{
		Name:     cfg.Name,
		IP:       cfg.IP,
		Port:     cfg.Port,
		Protocol: cfg.Protocol,
		IPType:   cfg.IPType,
	}

	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(
		cfg.RegistryAddr+"/refresh",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("successfully refreshed")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	errs := make(map[string]interface{})
	err = json.Unmarshal(body, &errs)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalf(
		"failed to refresh\ncode %d\nerrors: %v",
		resp.StatusCode,
		helpers.PrettyPrint(errs),
	)
}
