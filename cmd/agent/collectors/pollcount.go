package collectors

type PoolCount struct {
	count int64
}

func (p *PoolCount) GetMetrics() map[string]Value {
	var metrics = make(map[string]Value)
	p.count++
	metrics["PollCount"] = Value{VType: COUNTER, Value: p.count}
	return metrics
}
