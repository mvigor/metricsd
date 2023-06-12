package collectors

import "math/rand"

type Random struct {
}

func (r *Random) GetMetrics() map[string]Value {
	var metrics = make(map[string]Value)
	metrics["RandomValue"] = Value{VType: GAUGE, Value: rand.Float64()}
	return metrics
}
