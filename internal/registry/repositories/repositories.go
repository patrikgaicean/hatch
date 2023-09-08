package repositories

import (
	"github.com/patriuk/hatch/internal/registry/repositories/service"
	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	ServiceRepo service.ServiceRepository
}

type Params struct {
	Redis RedisDb
}

type RedisDb struct {
	Client *redis.Client
}

func SetupRepos(params Params) *Repositories {
	serviceRepo := service.NewServiceRepo(params.Redis.Client)

	return &Repositories{ServiceRepo: serviceRepo}
}
