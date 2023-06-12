package collectors

type PoolCount struct {
	count int64
}

func (p *PoolCount) GetMetrics() map[string]interface{} {
	var metrics = make(map[string]interface{})
	p.count++
	metrics["PollCount"] = p.count
	return metrics
}
