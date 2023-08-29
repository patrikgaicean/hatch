package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// todo: figure out if we still need the json tags
type Service struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IP        string `json:"ip"`
	Port      uint16 `json:"port"`
	Protocol  string `json:"protocol"`
	IPType    string `json:"ipType"`
	Address   string `json:"address"`
	Timestamp string `json:"timestamp"`
}

type ServiceRepository interface {
	RegisterService(service Service) error
	GetServiceByID(serviceID string) (Service, error)
	TestRedis()
	// Other service-related methods
}

type ServiceRepo struct {
	client *redis.Client
	ctx    context.Context
}

func NewRepo(client *redis.Client, ctx context.Context) ServiceRepository {
	return &ServiceRepo{client: client, ctx: ctx}
}

func (repo *ServiceRepo) TestRedis() {

	err := repo.client.Set(repo.ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := repo.client.Get(repo.ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)
}

// definitely change these to use hash instead of json
// we'll need to be able to query by service name
// actually, lemme list operations:

// register full body
// update timestamp
// delete by timestamp
// get all by name
// probably get by id, not sure if we can update without checking
// that the entry actually exists

func (repo *ServiceRepo) RegisterService(service Service) error {
	serviceJSON, err := json.Marshal(service)
	if err != nil {
		return err
	}
	err = repo.client.Set(repo.ctx, getServiceKey(service.ID), serviceJSON, 0).Err()
	return err
}

func (repo *ServiceRepo) GetServiceByID(serviceID string) (Service, error) {
	serviceJSON, err := repo.client.Get(repo.ctx, getServiceKey(serviceID)).Bytes()
	if err != nil {
		return Service{}, err
	}
	var service Service
	err = json.Unmarshal(serviceJSON, &service)
	if err != nil {
		return Service{}, err
	}
	return service, nil
}

func getServiceKey(serviceID string) string {
	return "service:" + serviceID
}
