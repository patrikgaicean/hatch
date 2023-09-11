package repositories

import (
	"context"
	"log"

	"github.com/patriuk/hatch/internal/common"
	"github.com/patriuk/hatch/internal/gateway/manager"
	"github.com/redis/go-redis/v9"
)

type ServiceRepository interface {
	GetAll() error
}

type ServiceRepo struct {
	client  *redis.Client
	manager *manager.Manager
}

func NewServiceRepo(rc *redis.Client, m *manager.Manager) ServiceRepository {
	return &ServiceRepo{client: rc, manager: m}
}

func (repo *ServiceRepo) GetAll() error {
	keys := common.ScanAllKeys(repo.client)
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

	// for _, v := range services {
	// 	fmt.Println(helpers.PrettyPrint(v))
	// }

	// store in manager
	repo.manager.UpdateInstances(services)

	return nil
}
