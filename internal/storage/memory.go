package storage

import (
	"fmt"
	"strconv"
)

type MetricRecord struct {
	VType VType
	Value interface{}
}

func (r MetricRecord) String() string {
	switch r.VType {
	case GAUGE:
		switch i := r.Value.(type) {
		case float64:
			return fmt.Sprintf("%g", i)
		case int:
			return fmt.Sprintf("%d", i)
		default:
			return ""
		}
	case COUNTER:
		switch i := r.Value.(type) {
		case int, int8, int16, int32, int64:
			return fmt.Sprintf("%d", i)
		default:
			return ""
		}
	}
	return "<unknown type>"
}

type MemoryStorage struct {
	Metrics map[string]MetricRecord
}

func (m *MemoryStorage) SetMetric(metricName string, metricValue string, metricType string) error {
	rec, err := m.ConvertData(metricValue, metricType)
	if err != nil {
		return err
	}
	if len(m.Metrics) == 0 {
		m.Metrics = make(map[string]MetricRecord)
	}
	m.Metrics[metricName] = rec
	return nil
}
func (m *MemoryStorage) GetMetric(metricName string) (MetricRecord, bool) {
	value, ok := m.Metrics[metricName]
	return value, ok
}
func (m *MemoryStorage) IndexMetrics() map[string]MetricRecord {
	return m.Metrics
}

func (m *MemoryStorage) ConvertData(metricValue string, metricType string) (MetricRecord, error) {
	switch metricType {
	case string(COUNTER):
		val, err := strconv.ParseInt(metricValue, 10, 64)

		if err != nil {
			return MetricRecord{}, err
		}
		return MetricRecord{VType: COUNTER, Value: val}, nil

	case string(GAUGE):
		val, err := strconv.ParseFloat(metricValue, 64)

		if err != nil {
			return MetricRecord{}, err
		}
		return MetricRecord{VType: GAUGE, Value: val}, nil
	default:
		return MetricRecord{}, fmt.Errorf("unkown type")
	}
}
