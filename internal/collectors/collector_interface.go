package collectors

import (
	"time"

	"github.com/mvigor/metricsd/internal/entities"
)

type Collector interface {
	GetMetrics() map[string]entities.MetricValue
	StartCollector(time.Duration) error
}
