package registry

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/patriuk/hatch/internal/registry/config"
	"github.com/patriuk/hatch/internal/registry/handlers"
	"github.com/patriuk/hatch/internal/registry/repositories"
	"github.com/patriuk/hatch/internal/registry/router"
	"github.com/patriuk/hatch/internal/registry/server"
	"github.com/redis/go-redis/v9"
)

func ListenAndServe(cfg config.Config) error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.IP, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	redisURL := fmt.Sprintf(
		"redis://%s:%s@%s:%d/",
		"default",
		cfg.Redis.Password,
		cfg.Redis.Host,
		cfg.Redis.Port,
	)

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	ctx := context.Background()
	repos := repositories.SetupRepos(repositories.Params{
		Redis: repositories.RedisDb{
			Client: client,
			Ctx:    ctx,
		},
	})

	handlers := handlers.SetupHandlers(*repos)
	router := router.SetupRoutes(*handlers)

	err = server.Serve(l, router)
	return err
}
