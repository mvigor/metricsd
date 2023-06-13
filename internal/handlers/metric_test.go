package handlers

import (
	"github.com/mvigor/metricsd/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowHandler(t *testing.T) {
	storage := storage.MemoryStorage{
		Metrics: map[string]storage.MetricRecord{
			"metric1": {
				VType: storage.GAUGE,
				Value: 0,
			},
			"metric2": {
				VType: storage.COUNTER,
				Value: 1,
			},
			"metric3": {
				VType: storage.GAUGE,
				Value: 444,
			},
			"metric4": {
				VType: storage.COUNTER,
				Value: 50000,
			},
		},
	}

	testCases := []struct {
		name                string
		method              string
		request             map[string]string
		expectedCode        int
		expectedBody        string
		expectedContentType string
	}{
		{
			name:                "test case #1",
			method:              http.MethodGet,
			request:             map[string]string{"metric_name": "metric1"},
			expectedCode:        http.StatusOK,
			expectedBody:        "0",
			expectedContentType: "text/html",
		},
		{
			name:                "test case #2",
			method:              http.MethodGet,
			request:             map[string]string{"metric_name": "metric3"},
			expectedCode:        http.StatusOK,
			expectedBody:        "444",
			expectedContentType: "text/html",
		},
		{
			name:                "test case #3",
			method:              http.MethodGet,
			request:             map[string]string{"metric_name": "metri"},
			expectedCode:        http.StatusNotFound,
			expectedBody:        "",
			expectedContentType: "text/html",
		},
		{
			name:         "test case #4",
			method:       http.MethodPut,
			request:      map[string]string{"metric_name": "metric12"},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "test case #5",
			method:       http.MethodDelete,
			request:      map[string]string{"metric_name": "metric12"},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "test case #6",
			method:       http.MethodPost,
			request:      map[string]string{"metric_name": "metric12"},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", nil)
			w := httptest.NewRecorder()
			handlerFunc := ShowHandler(tt.request, &storage)
			handlerFunc(w, request)
			res := w.Result()
			assert.Equal(t, res.StatusCode, tt.expectedCode)

			if res.StatusCode == http.StatusOK {
				defer res.Body.Close()
				resBody, err := io.ReadAll(res.Body)

				require.NoError(t, err)

				assert.Equal(t, tt.expectedBody, string(resBody))
				assert.Equal(t, tt.expectedContentType, res.Header.Get("Content-Type"))
			}
		})

	}
}

func TestUpdateHandler(t *testing.T) {
	type args struct {
		params  map[string]string
		storage storage.Storage
	}
	tests := []struct {
		name       string
		args       args
		method     string
		wantCount  int
		wantResult map[string]storage.MetricRecord
		wantStatus int
	}{
		{
			name:   "test case #1",
			method: http.MethodPost,
			args: args{
				params: map[string]string{
					"metric_type":  "counter",
					"metric_name":  "metric2",
					"metric_value": "102",
				},
				storage: &storage.MemoryStorage{
					Metrics: map[string]storage.MetricRecord{
						"metric1": {
							VType: storage.COUNTER,
							Value: 100,
						},
					},
				},
			},
			wantCount: 2,
			wantResult: map[string]storage.MetricRecord{
				"metric1": {
					VType: storage.COUNTER,
					Value: 100,
				},
				"metric2": {
					VType: storage.COUNTER,
					Value: 102,
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "test case #2",
			method: http.MethodPost,
			args: args{
				params: map[string]string{
					"metric_name":  "metric2",
					"metric_value": "102",
					"metric_type":  "gauge",
				},
				storage: &storage.MemoryStorage{
					Metrics: map[string]storage.MetricRecord{
						"metric2": {
							VType: storage.COUNTER,
							Value: 100,
						},
					},
				},
			},
			wantCount: 1,
			wantResult: map[string]storage.MetricRecord{
				"metric2": {
					VType: storage.COUNTER,
					Value: 102,
				},
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "test case #3",
			method: http.MethodPost,
			args: args{
				params: map[string]string{
					"metric2": "test",
				},
				storage: &storage.MemoryStorage{
					Metrics: map[string]storage.MetricRecord{
						"metric3": {
							VType: storage.GAUGE,
							Value: 0.0001,
						},
					},
				},
			},
			wantCount:  0,
			wantResult: map[string]storage.MetricRecord{},
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", nil)
			w := httptest.NewRecorder()
			handlerFunc := UpdateHandler(tt.args.params, tt.args.storage)
			handlerFunc(w, request)
			res := w.Result()
			assert.Equal(t, tt.wantStatus, res.StatusCode)
			defer res.Body.Close()
			if res.StatusCode == http.StatusOK {
				assert.Equal(t, len(tt.args.storage.IndexMetrics()), tt.wantCount)
			}
		})
	}

}
