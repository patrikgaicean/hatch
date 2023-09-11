package balancer

import (
	"errors"
	"sync"

	"github.com/patriuk/hatch/internal/common"
	"github.com/patriuk/hatch/internal/gateway/manager"
)

type RoundRobinBalancer struct {
	manager      *manager.Manager
	currentIndex map[string]int
	mu           sync.RWMutex
}

func NewRoundRobinBalancer(manager *manager.Manager) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		manager:      manager,
		currentIndex: make(map[string]int),
		mu:           sync.RWMutex{},
	}
}

func (rr *RoundRobinBalancer) GetNextInstance(serviceName string) (*common.Service, error) {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	instances := rr.manager.GetInstances(serviceName)
	if len(instances) == 0 {
		return nil, errors.New("no instances available")
	}

	instance := &instances[rr.currentIndex[serviceName]%len(instances)]
	rr.currentIndex[serviceName]++

	return instance, nil
}
