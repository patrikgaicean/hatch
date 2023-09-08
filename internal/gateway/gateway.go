package gateway

import (
	"github.com/patriuk/hatch/internal/gateway/config"
	"github.com/patriuk/hatch/internal/gateway/server"
)

func ListenAndServe(cfg config.Config) error {
	err := server.ListenAndServe(cfg.IP, cfg.Port)
	return err
}
