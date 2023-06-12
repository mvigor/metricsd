package storage

type Storage interface {
	SetMetric(metricName string, metricValue string) error
	GetMetric(metricName string) (string, bool)
	IndexMetrics() map[string]string
}

func GetStorage() Storage {
	return &MemoryStorage{}
}
