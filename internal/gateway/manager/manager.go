package manager

import (
	"sync"

	"github.com/patriuk/hatch/internal/common"
	"golang.org/x/exp/slices"
)

type Manager struct {
	services     map[string][]common.Service
	serviceNames []string
	mu           sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		services:     make(map[string][]common.Service),
		serviceNames: make([]string, 0),
		mu:           sync.RWMutex{},
	}
}

func (m *Manager) UpdateInstances(instances []common.Service) {
	services := make(map[string][]common.Service)

	for _, v := range instances {
		if _, ok := services[v.Name]; ok {
			services[v.Name] = append(services[v.Name], v)
		} else {
			services[v.Name] = []common.Service{v}
		}
	}

	serviceNames := make([]string, 0)
	for k := range services {
		serviceNames = append(serviceNames, k)
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	m.services = services
	m.serviceNames = serviceNames
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

func (m *Manager) HasService(name string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return slices.Contains(m.serviceNames, name)
}
