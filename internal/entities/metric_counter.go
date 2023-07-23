package entities

import "fmt"

type MetricCounter MetricValue

func NewMetricCounter(name string, val int64) (m *MetricCounter) {
	return &MetricCounter{
		Value: val,
		Name:  name,
		Type:  COUNTER,
	}
}

func (m *MetricCounter) SetValue(value MetricValue) {
	m.Value = m.Value.(int64) + value.Value.(int64)
}

func (m *MetricCounter) GetValue() interface{} {
	return m.Value
}

func (m *MetricCounter) GetStruct() MetricValue {
	return MetricValue{
		Value: m.Value,
		Name:  m.Name,
		Type:  m.Type,
	}
}

func (m *MetricCounter) ToString() string {
	return fmt.Sprintf("%d", m.Value)
}

func (m *MetricCounter) GetName() string {
	return m.Name
}

func (m *MetricCounter) GetHash() string {
	return string(m.Type) + "_" + m.Name
}
