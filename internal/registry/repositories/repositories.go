package repositories

import (
	"github.com/patriuk/hatch/internal/registry/repositories/service"
	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	ServiceRepo service.ServiceRepository
}

func SetupRepos(rc *redis.Client) *Repositories {
	serviceRepo := service.NewServiceRepo(rc)

	return &Repositories{ServiceRepo: serviceRepo}
}
