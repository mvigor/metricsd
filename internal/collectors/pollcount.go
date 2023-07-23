package collectors

import (
	"time"

	"github.com/mvigor/metricsd/internal/entities"
)

type PoolCount struct {
	count int64
}

func (p *PoolCount) GetMetrics() map[string]entities.MetricValue {
	var metrics = make(map[string]entities.MetricValue)
	p.count++
	metrics["PollCount"] = entities.MetricValue{Type: entities.COUNTER, Value: p.count}
	p.count = 0
	return metrics
}

func (p *PoolCount) StartCollector(poolInterval time.Duration) error {
	p.count = 0
	return nil
}
