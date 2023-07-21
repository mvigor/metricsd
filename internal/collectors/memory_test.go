package collectors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemory_GetMetrics(t *testing.T) {
	tests := []struct {
		name      string
		wantCount int
	}{
		{
			name:      "test case #1",
			wantCount: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memory{}
			m.StartCollector(1 * time.Second)
			time.Sleep(2 * time.Second)
			assert.Equalf(t, tt.wantCount, len(m.GetMetrics()), "GetMetrics()")
		})
	}
}
