package storage

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMemoryStorage_SetMetric(t *testing.T) {
	type fields struct {
		Metrics map[string]MetricRecord
	}
	type args struct {
		metricName  string
		metricValue string
		metricType  string
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
				Metrics: map[string]MetricRecord{},
			},
			args: args{
				metricName:  "metric 1",
				metricValue: "44",
				metricType:  "gauge",
			},
			wantErr:   false,
			wantCount: 1,
		},
		{
			name: "test case #2",
			fields: fields{
				Metrics: map[string]MetricRecord{
					"metric2": {
						Value: 35,
						VType: COUNTER,
					},
				},
			},
			args: args{
				metricName:  "metric2",
				metricValue: "33",
				metricType:  "counter",
			},
			wantErr:   false,
			wantCount: 1,
		},
		{
			name: "test case #3",
			fields: fields{
				Metrics: map[string]MetricRecord{
					"metric2": {
						Value: 35,
						VType: COUNTER,
					},
				},
			},
			args: args{
				metricName:  "metric3",
				metricValue: "33",
				metricType:  "unknown",
			},
			wantErr:   true,
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryStorage{
				Metrics: tt.fields.Metrics,
			}
			if tt.wantErr {
				assert.Error(t, m.SetMetric(tt.args.metricName, tt.args.metricValue, tt.args.metricType))
				return
			}
			assert.NoError(t, m.SetMetric(tt.args.metricName, tt.args.metricValue, tt.args.metricType))
			assert.Equal(t, len(m.Metrics), tt.wantCount)
		})
	}
}

func TestMemoryStorage_GetMetric(t *testing.T) {
	type fields struct {
		Metrics map[string]MetricRecord
	}
	type args struct {
		metricName  string
		metricValue interface{}
		metricType  VType
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOk    bool
		wantValue interface{}
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]MetricRecord{
					"metric1": {
						VType: COUNTER,
						Value: 11,
					},
					"metric2": {
						VType: GAUGE,
						Value: 44,
					},
					"metric3": {
						VType: COUNTER,
						Value: 55,
					},
					"metric4": {
						VType: COUNTER,
						Value: 66,
					},
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
				Metrics: map[string]MetricRecord{
					"metric1": {
						VType: COUNTER,
						Value: 33,
					},
					"metric2": {
						VType: GAUGE,
						Value: 44,
					},
					"metric3": {
						VType: COUNTER,
						Value: 55,
					},
					"metric4": {
						VType: COUNTER,
						Value: 66,
					},
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
				assert.Equal(t, tt.wantValue, v.Value)
			}
		})
	}
}

func TestMemoryStorage_IndexMetrics(t *testing.T) {
	type fields struct {
		Metrics map[string]MetricRecord
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]MetricRecord
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]MetricRecord{
					"metric1": {
						VType: COUNTER,
						Value: 55,
					},
					"metric2": {
						VType: COUNTER,
						Value: 66,
					},
					"metric3": {
						VType: COUNTER,
						Value: 77,
					},
					"metric4": {
						VType: COUNTER,
						Value: 88,
					},
				},
			},
			want: map[string]MetricRecord{
				"metric1": {
					VType: COUNTER,
					Value: 55,
				},
				"metric2": {
					VType: COUNTER,
					Value: 66,
				},
				"metric3": {
					VType: COUNTER,
					Value: 77,
				},
				"metric4": {
					VType: COUNTER,
					Value: 88,
				},
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
