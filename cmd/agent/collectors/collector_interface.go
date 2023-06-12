package collectors

type CounterType string

const (
	GAUGE   = CounterType("gauge")
	COUNTER = CounterType("counter")
)

type Value struct {
	VType CounterType
	Value interface{}
}
type Collector interface {
	GetMetrics() map[string]Value
}
