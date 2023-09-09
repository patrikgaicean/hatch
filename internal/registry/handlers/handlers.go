package handlers

import (
	"github.com/patriuk/hatch/internal/registry/handlers/discovery"
	"github.com/patriuk/hatch/internal/registry/repositories"
)

type Handlers struct {
	Discovery *discovery.Handler
}

func SetupHandlers(repos repositories.Repositories) *Handlers {
	return &Handlers{
		Discovery: discovery.NewHandler(repos.ServiceRepo),
	}
}
