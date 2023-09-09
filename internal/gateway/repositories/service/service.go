package service

import (
	"context"
	"fmt"
	"log"

	"github.com/patriuk/hatch/internal/common"
	"github.com/patriuk/hatch/internal/helpers"
	"github.com/redis/go-redis/v9"
)

type ServiceRepository interface {
	GetAll(name string) error
}

type ServiceRepo struct {
	client *redis.Client
}

func NewServiceRepo(rc *redis.Client) ServiceRepository {
	return &ServiceRepo{client: rc}
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
