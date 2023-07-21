package utils

import "github.com/mvigor/metricsd/internal/entities"

type MetricsMap map[string]entities.MetricValue

func (m1 MetricsMap) Add(m2 MetricsMap) {
	for k, v := range m2 {
		m1[k] = v
	}
}
