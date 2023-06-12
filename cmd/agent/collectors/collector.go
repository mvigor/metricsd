package collectors

var collectors = []Collector{
	&Memory{},
	&PoolCount{},
	&Random{},
}

func CollectData() map[string]interface{} {
	metrics := make(map[string]interface{})
	for _, collector := range collectors {
		res := collector.GetMetrics()
		mergeMaps(metrics, res)
	}
	return metrics
}

func mergeMaps(m1 map[string]interface{}, m2 map[string]interface{}) {
	for k, v := range m2 {
		m1[k] = v
	}
}
