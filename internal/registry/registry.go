package registry

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/patriuk/hatch/internal/registry/config"
	"github.com/patriuk/hatch/internal/registry/handlers"
	"github.com/patriuk/hatch/internal/registry/repositories"
	"github.com/patriuk/hatch/internal/registry/router"
	"github.com/patriuk/hatch/internal/registry/server"
	"github.com/redis/go-redis/v9"
)

func ListenAndServe(cfg config.Config) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.IP, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

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
	repos := repositories.SetupRepos(redisClient)
	handlers := handlers.SetupHandlers(*repos)
	router := router.SetupRoutes(*handlers)

	// cleanup every x seconds based on cfg.Cleanup
	ticker := time.NewTicker(time.Duration(cfg.Cleanup) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				repos.ServiceRepo.Cleanup(cfg.Cleanup)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	// close(quit) on server graceful shutdown ? though it will get closed
	// if server crashes anyway..

	err = server.Serve(listener, router)
	return err
}
