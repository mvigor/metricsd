package entities

import (
	"errors"
	"strconv"
	"strings"
)

type MetricType string

const (
	GAUGE   = MetricType("gauge")
	COUNTER = MetricType("counter")
)

type Metric interface {
	SetValue(MetricValue)
	GetValue() interface{}
	GetStruct() MetricValue
	ToString() string
	GetHash() string
	GetName() string
}

type MetricValue struct {
	Name  string
	Type  MetricType
	Value interface{}
}

func MetricFactory(metricType string, metricName string, metricValue string) (Metric, error) {
	metricType = strings.ToLower(metricType)
	switch metricType {
	case "gauge":
		val, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return nil, err
		}
		return NewMetricGauge(metricName, val), nil
	case "counter":
		val, err := strconv.Atoi(metricValue)
		if err != nil {
			return nil, err
		}
		return NewMetricCounter(metricName, val), nil
	}
	return nil, errors.New("invalid metric type")
}
