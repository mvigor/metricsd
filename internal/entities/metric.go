package entities

type MetricType string

const (
	GAUGE   = MetricType("gauge")
	COUNTER = MetricType("counter")
)

type Metric interface {
	SetValue(interface{})
	GetValue() interface{}
	ToString() string
}

type MetricValue struct {
	Type  MetricType
	Value interface{}
}

func NewMetricValue(metricType MetricType) *MetricValue {
	return &MetricValue{
		Type: metricType,
	}
}
