package collectors

import "math/rand"

type Random struct {
}

func (r *Random) GetMetrics() map[string]interface{} {
	var metrics = make(map[string]interface{})
	metrics["RandomValue"] = rand.Float64()
	return metrics
}
