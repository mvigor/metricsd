package storage

type MemoryStorage struct {
	Metrics map[string]float64
}

func (m *MemoryStorage) SetMetric(metricName string, metricValue float64) error {
	m.Metrics[metricName] = metricValue
	return nil
}
func (m *MemoryStorage) GetMetric(metricName string) (float64, bool) {
	value, ok := m.Metrics[metricName]
	return value, ok
}
func (m *MemoryStorage) IndexMetrics() map[string]float64 {
	return m.Metrics
}
