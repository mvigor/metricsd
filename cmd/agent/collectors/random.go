package collectors

import (
	"math/rand"
	"time"
)

type Random struct {
}

func (r *Random) GetMetrics() map[string]Value {
	var metrics = make(map[string]Value)
	metrics["RandomValue"] = Value{VType: GAUGE, Value: rand.Float64()}
	return metrics
}

func (r *Random) StartCollector(poolInterval time.Duration) error {
	return nil
}
