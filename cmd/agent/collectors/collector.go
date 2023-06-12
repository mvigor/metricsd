package collectors

import (
	"fmt"
	"time"
)

var collectors = []Collector{
	&Memory{},
	&PoolCount{},
	&Random{},
}

func CollectData() map[string]Value {
	metrics := make(map[string]Value)
	for _, collector := range collectors {
		res := collector.GetMetrics()
		mergeMaps(metrics, res)
	}
	return metrics
}

func StartCollectors(poolInterval time.Duration) {
	for _, collector := range collectors {
		err := collector.StartCollector(poolInterval)
		if err != nil {
			panic(fmt.Sprintf("couldn't initialize collector %T", collector))
		}
	}
}

func mergeMaps(m1 map[string]Value, m2 map[string]Value) {
	for k, v := range m2 {
		m1[k] = v
	}
}
