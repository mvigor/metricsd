package storage

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMemoryStorage_SetMetric(t *testing.T) {
	type fields struct {
		Metrics map[string]string
	}
	type args struct {
		metricName  string
		metricValue string
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
				Metrics: map[string]string{},
			},
			args: args{
				metricName:  "metric 1",
				metricValue: "value1",
			},
			wantErr:   false,
			wantCount: 1,
		},
		{
			name: "test case #2",
			fields: fields{
				Metrics: map[string]string{
					"metric2": "value2",
				},
			},
			args: args{
				metricName:  "metric2",
				metricValue: "value3",
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
		Metrics map[string]string
	}
	type args struct {
		metricName  string
		metricValue string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOk    bool
		wantValue string
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]string{
					"metric1": "value1",
					"metric2": "value2",
					"metric3": "value3",
					"metric4": "value4",
				},
			},
			args: args{
				metricName: "metric1",
			},
			wantOk:    true,
			wantValue: "value1",
		},
		{
			name: "test case #2",
			fields: fields{
				Metrics: map[string]string{
					"metric1": "value1",
					"metric2": "value2",
					"metric3": "value3",
					"metric4": "value4",
				},
			},
			args: args{
				metricName:  "metric5",
				metricValue: "value3",
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
		Metrics map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]string{
					"metric1": "value1",
					"metric2": "value2",
					"metric3": "value3",
					"metric4": "value4",
				},
			},
			want: map[string]string{
				"metric1": "value1",
				"metric2": "value2",
				"metric3": "value3",
				"metric4": "value4",
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
