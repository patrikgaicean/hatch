package gateway

import (
	"fmt"

	"github.com/patriuk/hatch/internal/gateway/config"
	"github.com/patriuk/hatch/internal/gateway/repositories"
	"github.com/patriuk/hatch/internal/gateway/server"
	"github.com/redis/go-redis/v9"
)

func ListenAndServe(cfg config.Config) error {
	URL := fmt.Sprintf(
		"redis://%s:%s@%s:%d/",
		"default",
		cfg.Redis.Password,
		cfg.Redis.Host,
		cfg.Redis.Port,
	)

	opt, err := redis.ParseURL(URL)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(opt)
	// repo := repositories.SetupRepos(redisClient)
	repositories.SetupRepos(redisClient)

	err = server.ListenAndServe(cfg.IP, cfg.Port)
	return err
}
