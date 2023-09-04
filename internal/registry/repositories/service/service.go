package service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/patriuk/hatch/internal/helpers"
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

type ServiceRepository interface {
	Register(service Service) error
	Unregister(service Service) error
	Refresh(service Service) error
	GetAllByName(name string) error
}

type ServiceRepo struct {
	client *redis.Client
	ctx    context.Context
}

func NewServiceRepo(client *redis.Client, ctx context.Context) ServiceRepository {
	return &ServiceRepo{client: client, ctx: ctx}
}

func (repo *ServiceRepo) Register(service Service) error {
	key := getServiceKey(service)

	err := repo.client.HSet(repo.ctx, key, service).Err()
	if err != nil {
		fmt.Println("RegisterService error")
	}

	return nil
}

func (repo *ServiceRepo) Unregister(service Service) error {
	key := getServiceKey(service)

	err := repo.client.HDel(repo.ctx, key).Err()
	if err != nil {
		fmt.Println("UnregisterService error")
	}

	return nil
}

func (repo *ServiceRepo) Refresh(service Service) error {
	fmt.Println("in refresh")
	key := getServiceKey(service)

	res, err := repo.client.HSet(repo.ctx, key, "timestamp", service.Timestamp).Result()
	if err != nil {
		fmt.Println("RegisterService error")
	}

	if res == 0 {
		fmt.Printf("Key %s does not exist", key)
	} else {
		fmt.Printf("Updated %s key with timestamp", key)
	}

	return nil
}

func (repo *ServiceRepo) GetAllByName(name string) error {
	pattern := fmt.Sprintf("%s:*", name)
	fmt.Println("pattern - ", pattern)

	var cursor uint64
	var keys []string

	for {
		var foundKeys []string
		var err error
		foundKeys, cursor, err = repo.client.Scan(context.Background(), cursor, pattern, 0).Result()
		if err != nil {
			log.Fatal(err)
		}
		keys = append(keys, foundKeys...)

		if cursor == 0 {
			break
		}
	}

	var services []Service
	for _, key := range keys {
		val, err := repo.client.HGetAll(repo.ctx, key).Result()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Key:", key)
		s := hashToModel(val)
		services = append(services, s)
	}

	for _, v := range services {
		fmt.Println(helpers.PrettyPrint(v))
	}

	return nil
}

func getServiceKey(service Service) string {
	serialized, err := json.Marshal(ServiceHash{
		Name:     service.Name,
		IP:       service.IP,
		Port:     service.Port,
		Protocol: service.Protocol,
		IPType:   service.IPType,
	})
	if err != nil {
		// handle error
	}

	hash := sha256.Sum256(serialized)
	hashString := fmt.Sprintf("%x", hash)
	key := fmt.Sprintf("%s:%s", service.Name, hashString)

	return key
}

func hashToModel(hashData map[string]string) Service {
	var service Service

	elem := reflect.ValueOf(&service).Elem()
	typeOfElem := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		tag := typeOfElem.Field(i).Tag.Get("redis")

		if value, ok := hashData[tag]; ok {
			switch field.Type().Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Uint16:
				var parsedValue uint16
				fmt.Sscanf(value, "%d", &parsedValue)
				field.SetUint(uint64(parsedValue))
			case reflect.Int64:
				var parsedValue int64
				fmt.Sscanf(value, "%d", &parsedValue)
				field.SetInt(parsedValue)
			}
		}
	}

	return service
}
