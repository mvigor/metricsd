package collectors

import (
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
		want   map[string]interface{}
	}{
		{
			name: "test case #1",
			fields: fields{
				count: 0,
			},
			want: map[string]interface{}{
				"PollCount": int64(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PoolCount{
				count: tt.fields.count,
			}
			assert.Equal(t, tt.want, p.GetMetrics(), "GetMetrics()")
		})
	}
}
