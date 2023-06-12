package storage

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMemoryStorage_SetMetric(t *testing.T) {
	type fields struct {
		Metrics map[string]float64
	}
	type args struct {
		metricName  string
		metricValue float64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantCount int
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]float64{},
			},
			args: args{
				metricName:  "metric 1",
				metricValue: 44,
			},
			wantErr:   false,
			wantCount: 1,
		},
		{
			name: "test case #2",
			fields: fields{
				Metrics: map[string]float64{
					"metric2": 4,
				},
			},
			args: args{
				metricName:  "metric2",
				metricValue: 33,
			},
			wantErr:   false,
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryStorage{
				Metrics: tt.fields.Metrics,
			}
			assert.NoError(t, m.SetMetric(tt.args.metricName, tt.args.metricValue))
			assert.Equal(t, len(m.Metrics), tt.wantCount)
		})
	}
}

func TestMemoryStorage_GetMetric(t *testing.T) {
	type fields struct {
		Metrics map[string]float64
	}
	type args struct {
		metricName  string
		metricValue float64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOk    bool
		wantValue float64
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]float64{
					"metric1": 11,
					"metric2": 33,
					"metric3": 55,
					"metric4": 66,
				},
			},
			args: args{
				metricName: "metric1",
			},
			wantOk:    true,
			wantValue: 11,
		},
		{
			name: "test case #2",
			fields: fields{
				Metrics: map[string]float64{
					"metric1": 101,
					"metric2": 102,
					"metric3": 103,
					"metric4": 104,
				},
			},
			args: args{
				metricName:  "metric5",
				metricValue: 11,
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryStorage{
				Metrics: tt.fields.Metrics,
			}
			v, ok := m.GetMetric(tt.args.metricName)
			assert.Equal(t, tt.wantOk, ok)
			if ok {
				assert.Equal(t, tt.wantValue, v)
			}
		})
	}
}

func TestMemoryStorage_IndexMetrics(t *testing.T) {
	type fields struct {
		Metrics map[string]float64
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]float64
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]float64{
					"metric1": 55,
					"metric2": 66,
					"metric3": 77,
					"metric4": 88,
				},
			},
			want: map[string]float64{
				"metric1": 55,
				"metric2": 66,
				"metric3": 77,
				"metric4": 88,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryStorage{
				Metrics: tt.fields.Metrics,
			}
			assert.True(t, reflect.DeepEqual(tt.want, m.IndexMetrics()), "IndexMetrics()")
		})
	}
}
