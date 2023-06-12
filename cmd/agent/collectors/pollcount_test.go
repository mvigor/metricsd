package collectors

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoolCount_GetMetrics(t *testing.T) {
	type fields struct {
		count int64
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]Value
	}{
		{
			name: "test case #1",
			fields: fields{
				count: 0,
			},
			want: map[string]Value{
				"PollCount": {
					VType: COUNTER,
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
