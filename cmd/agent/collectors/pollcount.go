package collectors

import "time"

type PoolCount struct {
	count int64
}

func (p *PoolCount) GetMetrics() map[string]Value {
	var metrics = make(map[string]Value)
	p.count++
	metrics["PollCount"] = Value{VType: COUNTER, Value: p.count}
	return metrics
}

func (p *PoolCount) StartCollector(poolInterval time.Duration) error {
	p.count = 0
	return nil
}
