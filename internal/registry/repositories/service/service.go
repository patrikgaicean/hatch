package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/patriuk/hatch/internal/common"
	"github.com/patriuk/hatch/internal/helpers"
	"github.com/redis/go-redis/v9"
)

type ServiceRepository interface {
	Register(s common.Service) error
	Unregister(s common.Service) error
	Refresh(s common.Service) error
	GetAll(name string) error
	Cleanup(ttl int64) error
}

type ServiceRepo struct {
	client *redis.Client
}

func NewServiceRepo(rc *redis.Client) ServiceRepository {
	return &ServiceRepo{client: rc}
}

func (repo *ServiceRepo) Register(s common.Service) error {
	key := common.GetServiceKey(s)

	ctx := context.Background()
	err := repo.client.HSet(ctx, key, s).Err()

	if err != nil {
		fmt.Println("RegisterService error")
	}

	return nil
}

func (repo *ServiceRepo) Unregister(s common.Service) error {
	key := common.GetServiceKey(s)

	ctx := context.Background()
	err := repo.client.Del(ctx, key).Err()
	if err != nil {
		fmt.Println("UnregisterService error")
	}

	return nil
}

func (repo *ServiceRepo) Refresh(s common.Service) error {
	key := common.GetServiceKey(s)

	ctx := context.Background()
	_, err := repo.client.HSet(ctx, key, "timestamp", s.Timestamp).Result()
	if err != nil {
		fmt.Println("Error updating timestamp:", err)
		return err
	}

	fmt.Printf("Updated %s key with timestamp %d\n", key, s.Timestamp)

	return nil
}

func (repo *ServiceRepo) GetAll(name string) error {
	pattern := ""
	if len(name) != 0 {
		pattern = fmt.Sprintf("%s:*", name)
	}
	keys := common.ScanAllKeys(repo.client, pattern)
	ctx := context.Background()

	var services []common.Service
	for _, key := range keys {
		var s common.Service
		err := repo.client.HGetAll(ctx, key).Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		services = append(services, s)
	}

	for _, v := range services {
		fmt.Println(helpers.PrettyPrint(v))
	}

	return nil
}

func (repo *ServiceRepo) Cleanup(ttl int64) error {
	keys := common.ScanAllKeys(repo.client, "")
	ctx := context.Background()

	for _, key := range keys {
		val, err := repo.client.HGetAll(ctx, key).Result()
		if err != nil {
			log.Fatal(err)
		}

		t := time.Now().Unix()
		ts, err := strconv.Atoi(val["timestamp"])
		if err != nil {
			log.Fatal(err)
		}
		ti := int64(ts)

		if t-ti > ttl {
			err := repo.client.Del(ctx, key).Err()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
