package storage

import "github.com/mvigor/metricsd/internal/entities"

type VType string

var stor = &MemoryStorage{
	Metrics: map[string]entities.Metric{},
}

const (
	GAUGE   = VType("gauge")
	COUNTER = VType("counter")
)

type Storage interface {
	SetMetric(metric entities.Metric) error
	GetMetric(metricType string, metricName string) (entities.Metric, bool)
	IndexMetrics() map[string]entities.Metric
}

func GetStorage() Storage {
	return stor
}
