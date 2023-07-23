package collectors

import (
	"math/rand"
	"time"

	"github.com/mvigor/metricsd/internal/entities"
)

type Random struct {
}

func (r *Random) GetMetrics() map[string]entities.MetricValue {
	var metrics = make(map[string]entities.MetricValue)
	metrics["RandomValue"] = entities.MetricValue{Type: entities.GAUGE, Value: rand.Float64()}
	return metrics
}

func (r *Random) StartCollector(poolInterval time.Duration) error {
	return nil
}
