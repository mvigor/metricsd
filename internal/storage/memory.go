package storage

import (
	"github.com/mvigor/metricsd/internal/entities"
)

type MemoryStorage struct {
	Metrics map[string]entities.Metric
}

func (m *MemoryStorage) SetMetric(metric entities.Metric) error {
	existMetric, ok := m.Metrics[metric.GetHash()]
	if !ok {
		m.Metrics[metric.GetHash()] = metric
		return nil
	}
	existMetric.SetValue(metric.GetStruct())
	return nil
}
func (m *MemoryStorage) GetMetric(metricType string, metricName string) (entities.Metric, bool) {
	value, ok := m.Metrics[metricType+"_"+metricName]
	return value, ok
}
func (m *MemoryStorage) IndexMetrics() map[string]entities.Metric {
	return m.Metrics
}
