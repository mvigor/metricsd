package collectors

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCollectData(t *testing.T) {
	tests := []struct {
		name      string
		wantCount int
	}{
		{
			name:      "test case #1",
			wantCount: 29,
		},
	}
	StartCollectors(1 * time.Second)
	time.Sleep(2 * time.Second)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := CollectData()
			assert.Equal(t, tt.wantCount, len(collected))
		})
	}
}
