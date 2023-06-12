package collectors

import (
	"encoding/json"
	"runtime"
)

var memoryMetrics = map[string]bool{
	"Alloc":         true,
	"BuckHashSys":   true,
	"Frees":         true,
	"GCCPUFraction": true,
	"GCSys":         true,
	"HeapAlloc":     true,
	"HeapIdle":      true,
	"HeapInuse":     true,
	"HeapObjects":   true,
	"HeapReleased":  true,
	"HeapSys":       true,
	"LastGC":        true,
	"Lookups":       true,
	"MCacheInuse":   true,
	"MCacheSys":     true,
	"MSpanInuse":    true,
	"MSpanSys":      true,
	"Mallocs":       true,
	"NextGC":        true,
	"NumForcedGC":   true,
	"NumGC":         true,
	"OtherSys":      true,
	"PauseTotalNs":  true,
	"StackInuse":    true,
	"StackSys":      true,
	"Sys":           true,
	"TotalAlloc":    true,
}

type Memory struct {
}

func (m *Memory) GetMetrics() map[string]Value {
	var rtm runtime.MemStats
	var metrics = make(map[string]Value)
	runtime.ReadMemStats(&rtm)
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(rtm)

	err := json.Unmarshal(inrec, &inInterface)
	if err != nil {
		panic("invalid struct")
	}

	for field, val := range inInterface {
		_, ok := memoryMetrics[field]
		if ok {
			metrics[field] = Value{VType: GAUGE, Value: val}
		}
	}

	return metrics
}
