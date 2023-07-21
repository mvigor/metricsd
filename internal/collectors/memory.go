package collectors

import (
	"encoding/json"
	"runtime"
	"time"

	entities "github.com/mvigor/metricsd/internal/entities"
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
	values map[string]entities.MetricValue
}

func (m *Memory) GetMetrics() map[string]entities.MetricValue {
	return m.values
}

func (m *Memory) StartCollector(poolInterval time.Duration) error {
	m.values = make(map[string]entities.MetricValue)
	go func() {

		for {
			var rtm runtime.MemStats
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
					m.values[field] = entities.MetricValue{Type: entities.GAUGE, Value: val}
				}
			}

			time.Sleep(poolInterval)
		}
	}()
	return nil
}
