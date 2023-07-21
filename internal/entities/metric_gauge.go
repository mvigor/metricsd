package entities

import "fmt"

type MetricGauge struct {
	MetricValue
}

func (m *MetricGauge) SetValue(val interface{}) {
	m.Value = val
	m.Type = GAUGE
}

func (m *MetricGauge) GetValue() interface{} {
	return m.Value
}

func (m *MetricGauge) ToString() string {
	return fmt.Sprintf("%g", m.Value)
}
