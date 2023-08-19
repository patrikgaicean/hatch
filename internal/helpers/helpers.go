package helpers

import (
	"encoding/json"
	"fmt"
	"net"
)

func GetIPType(ip string) (string, error) {
	pip := net.ParseIP(ip)
	if pip == nil {
		// fmt.Printf("%s is not a valid IP address.\n", ip)
		return "", fmt.Errorf("%s is not a valid IP address.\n", ip)
	} else if pip.To4() != nil {
		// fmt.Printf("%s is IPv4.\n", ip)
		return "IPv4", nil
	} else {
		// fmt.Printf("%s is IPv6.\n", ip)
		return "IPv6", nil
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
