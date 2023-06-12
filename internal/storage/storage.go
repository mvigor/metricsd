package storage

type VType string

const (
	GAUGE   = VType("gauge")
	COUNTER = VType("counter")
)

type Storage interface {
	SetMetric(metricName string, metricValue string, metricType string) error
	GetMetric(metricName string) (MetricRecord, bool)
	IndexMetrics() map[string]MetricRecord
}

func GetStorage() Storage {
	return &MemoryStorage{}
}
