package repositories

import (
	"context"

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
	Ctx    context.Context
}

func SetupRepos(params Params) *Repositories {
	serviceRepo := service.NewServiceRepo(
		params.Redis.Client,
		params.Redis.Ctx,
	)

	return &Repositories{ServiceRepo: serviceRepo}
}
