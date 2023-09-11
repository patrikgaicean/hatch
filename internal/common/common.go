package common

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	Name      string `redis:"name"`
	IP        string `redis:"ip"`
	Port      uint16 `redis:"port"`
	Protocol  string `redis:"protocol"`
	IPType    string `redis:"ipType"`
	Timestamp int64  `redis:"timestamp"`
}

type ServiceHash struct {
	Name     string
	IP       string
	Port     uint16
	Protocol string
	IPType   string
}

func GetServiceKey(s Service) string {
	serialized, err := json.Marshal(ServiceHash{
		Name:     s.Name,
		IP:       s.IP,
		Port:     s.Port,
		Protocol: s.Protocol,
		IPType:   s.IPType,
	})
	if err != nil {
		// handle error
	}

	hash := sha256.Sum256(serialized)
	hashString := fmt.Sprintf("%x", hash)
	key := fmt.Sprintf("%s:%s", s.Name, hashString)

	return key
}

func ScanAllKeys(rc *redis.Client) []string {
	var keys []string
	var cursor uint64
	ctx := context.Background()

	for {
		var foundKeys []string
		var err error
		foundKeys, cursor, err = rc.Scan(ctx, cursor, "", 0).Result()
		if err != nil {
			log.Fatal(err)
		}
		keys = append(keys, foundKeys...)

		if cursor == 0 {
			break
		}
	}

	return keys
}
