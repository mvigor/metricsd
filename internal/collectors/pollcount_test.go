package collectors

import (
	"encoding/json"
	"testing"

	"github.com/mvigor/metricsd/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestPoolCount_GetMetrics(t *testing.T) {
	type fields struct {
		count int64
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]entities.MetricValue
	}{
		{
			name: "test case #1",
			fields: fields{
				count: 0,
			},
			want: map[string]entities.MetricValue{
				"PollCount": {
					Type:  entities.COUNTER,
					Value: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PoolCount{
				count: tt.fields.count,
			}
			expect, _ := json.Marshal(tt.want)
			actual, _ := json.Marshal(p.GetMetrics())
			assert.Equal(t, expect, actual, "GetMetrics()")
		})
	}
}
