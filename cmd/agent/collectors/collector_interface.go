package collectors

type Collector interface {
	GetMetrics() map[string]interface{}
}
