package entities

import "fmt"

type MetricCounter struct {
	MetricValue
}

func (m *MetricCounter) SetValue(val interface{}) {
	m.Value = m.Value.(int64) + val.(int64)
	m.Type = COUNTER
}

func (m *MetricCounter) GetValue() interface{} {
	return m.Value
}

func (m *MetricCounter) ToString() string {
	return fmt.Sprintf("%d", m.Value)
}
