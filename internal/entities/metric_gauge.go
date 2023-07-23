package entities

import (
	"fmt"
)

type MetricGauge MetricValue

func NewMetricGauge(name string, val float64) (m *MetricGauge) {
	return &MetricGauge{
		Value: val,
		Name:  name,
		Type:  GAUGE,
	}
}

func (m *MetricGauge) SetValue(val MetricValue) {
	m.Value = val.Value
}

func (m *MetricGauge) GetStruct() MetricValue {
	return MetricValue{
		Value: m.Value,
		Name:  m.Name,
		Type:  m.Type,
	}
}

func (m *MetricGauge) GetValue() interface{} {
	return m.Value
}

func (m *MetricGauge) GetName() string {
	return m.Name
}

func (m *MetricGauge) ToString() string {
	return fmt.Sprintf("%g", m.Value)
}

func (m *MetricGauge) GetHash() string {
	return string(m.Type) + "_" + m.Name
}
