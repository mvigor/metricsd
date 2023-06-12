package collectors

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
			assert.Equalf(t, tt.wantCount, len(m.GetMetrics()), "GetMetrics()")
		})
	}
}
