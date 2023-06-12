package storage

type MemoryStorage struct {
	Metrics map[string]string
}

func (m *MemoryStorage) SetMetric(metricName string, metricValue string) error {
	m.Metrics[metricName] = metricValue
	return nil
}
func (m *MemoryStorage) GetMetric(metricName string) (string, bool) {
	value, ok := m.Metrics[metricName]
	return value, ok
}
func (m *MemoryStorage) IndexMetrics() map[string]string {
	return m.Metrics
}
