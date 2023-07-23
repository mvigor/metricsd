package storage

import (
	"reflect"
	"testing"

	"github.com/mvigor/metricsd/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStorage_SetMetric(t *testing.T) {
	type fields struct {
		Metrics map[string]entities.Metric
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
				Metrics: map[string]entities.Metric{},
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
				Metrics: map[string]entities.Metric{
					"counter_metric2": entities.NewMetricCounter("metric2", 0),
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
				Metrics: map[string]entities.Metric{
					"counter_metric2": entities.NewMetricCounter("metric2", 35),
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
			metric, err := entities.MetricFactory(tt.args.metricType, tt.args.metricName, tt.args.metricValue)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			m := &MemoryStorage{
				Metrics: tt.fields.Metrics,
			}
			if tt.wantErr {
				assert.Error(t, m.SetMetric(metric))
				return
			}
			assert.NoError(t, m.SetMetric(metric))
			assert.Equal(t, len(m.Metrics), tt.wantCount)
		})
	}
}

func TestMemoryStorage_GetMetric(t *testing.T) {
	type fields struct {
		Metrics map[string]entities.Metric
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
				Metrics: map[string]entities.Metric{
					"counter_metric1": entities.NewMetricCounter("metric1", 11),
					"gauge_metric2":   entities.NewMetricGauge("metric2", 44),
					"counter_metric3": entities.NewMetricCounter("metric3", 55),
					"counter_metric4": entities.NewMetricCounter("metric5", 66),
				},
			},
			args: args{
				metricName: "metric1",
				metricType: COUNTER,
			},
			wantOk:    true,
			wantValue: 11,
		},
		{
			name: "test case #2",
			fields: fields{
				Metrics: map[string]entities.Metric{
					"counter_metric1": entities.NewMetricCounter("metric1", 33),
					"gauge_metric2":   entities.NewMetricGauge("metric2", 44),
					"counter_metric3": entities.NewMetricCounter("metric3", 55),
					"counter_metric4": entities.NewMetricCounter("metric4", 66),
				},
			},
			args: args{
				metricName:  "metric5",
				metricValue: 11,
				metricType:  COUNTER,
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryStorage{
				Metrics: tt.fields.Metrics,
			}
			v, ok := m.GetMetric(string(tt.args.metricType), tt.args.metricName)
			assert.Equal(t, tt.wantOk, ok)
			if ok {
				assert.Equal(t, tt.wantValue, v.GetValue())
			}
		})
	}
}

func TestMemoryStorage_IndexMetrics(t *testing.T) {
	type fields struct {
		Metrics map[string]entities.Metric
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]entities.Metric
	}{
		{
			name: "test case #1",
			fields: fields{
				Metrics: map[string]entities.Metric{
					"counter_metric1": entities.NewMetricCounter("metric1", 55),
					"counter_metric2": entities.NewMetricCounter("metric2", 66),
					"counter_metric3": entities.NewMetricCounter("metric3", 77),
					"counter_metric4": entities.NewMetricCounter("metric4", 88),
				},
			},
			want: map[string]entities.Metric{
				"counter_metric1": entities.NewMetricCounter("metric1", 55),
				"counter_metric2": entities.NewMetricCounter("metric2", 66),
				"counter_metric3": entities.NewMetricCounter("metric3", 77),
				"counter_metric4": entities.NewMetricCounter("metric4", 88),
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
