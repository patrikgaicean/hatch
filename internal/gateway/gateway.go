package gateway

import (
	"fmt"
	"time"

	"github.com/patriuk/hatch/internal/gateway/balancer"
	"github.com/patriuk/hatch/internal/gateway/config"
	"github.com/patriuk/hatch/internal/gateway/handlers"
	"github.com/patriuk/hatch/internal/gateway/manager"
	"github.com/patriuk/hatch/internal/gateway/repositories"
	"github.com/patriuk/hatch/internal/gateway/router"
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

	manager := manager.NewManager()
	balancer := balancer.NewRoundRobinBalancer(manager)

	opt, err := redis.ParseURL(URL)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(opt)
	serviceRepo := repositories.NewServiceRepo(redisClient, manager)

	handler := handlers.NewGatewayHandler(balancer, serviceRepo, manager)
	router := router.SetupRoutes(handler)

	// setup initial services
	serviceRepo.GetAll()

	// need a config for this duration
	ticker := time.NewTicker(time.Duration(60) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				serviceRepo.GetAll()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	err = server.ListenAndServe(cfg.IP, cfg.Port, router)
	return err
}
