package manager

import (
	"sync"

	"github.com/patriuk/hatch/internal/common"
)

type Manager struct {
	services map[string][]common.Service
	mu       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		services: make(map[string][]common.Service),
		mu:       sync.RWMutex{},
	}
}

func (m *Manager) UpdateInstances(services []common.Service) {
	instances := make(map[string][]common.Service)

	for _, v := range services {
		if _, ok := instances[v.Name]; ok {
			instances[v.Name] = append(instances[v.Name], v)
		} else {
			instances[v.Name] = []common.Service{v}
		}
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	m.services = instances
}

func (m *Manager) GetInstances(name string) []common.Service {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.services[name]
}

func (m *Manager) GetAllInstances() map[string][]common.Service {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.services
}
