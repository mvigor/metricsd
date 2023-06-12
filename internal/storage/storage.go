package storage

type Storage interface {
	SetMetric(metricName string, metricValue float64) error
	GetMetric(metricName string) (float64, bool)
	IndexMetrics() map[string]float64
}

func GetStorage() Storage {
	return &MemoryStorage{}
}
