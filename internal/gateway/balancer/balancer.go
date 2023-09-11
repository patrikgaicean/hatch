package balancer

import (
	"github.com/patriuk/hatch/internal/common"
)

type Balancer interface {
	GetNextInstance(serviceName string) (*common.Service, error)
}
