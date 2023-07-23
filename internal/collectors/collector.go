package collectors

import (
	"fmt"
	"time"

	"github.com/mvigor/metricsd/internal/entities"
)

var collectors = []Collector{
	&Memory{},
	&PoolCount{},
	&Random{},
}

func CollectData() map[string]entities.MetricValue {
	metrics := make(map[string]entities.MetricValue)
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

func mergeMaps(m1 map[string]entities.MetricValue, m2 map[string]entities.MetricValue) {
	for k, v := range m2 {
		m1[k] = v
	}
}
