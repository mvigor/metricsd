package handlers

import (
	"github.com/mvigor/metricsd/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestShowHandler(t *testing.T) {
	storage := storage.MemoryStorage{
		Metrics: map[string]float64{
			"metric1": 0,
			"metric2": 1,
			"metric3": 444,
			"metric4": 444,
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
			expectedBody:        "metric name = metric1, value = 0",
			expectedContentType: "text/html",
		},
		{
			name:                "test case #2",
			method:              http.MethodGet,
			request:             map[string]string{"metric_name": "metric3"},
			expectedCode:        http.StatusOK,
			expectedBody:        "metric name = metric3, value = 444",
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
		wantResult map[string]float64
		wantStatus int
	}{
		{
			name:   "test case #1",
			method: http.MethodPost,
			args: args{
				params: map[string]string{
					"metric_name":  "metric2",
					"metric_value": "102",
				},
				storage: &storage.MemoryStorage{
					Metrics: map[string]float64{
						"metric1": 100,
					},
				},
			},
			wantCount: 2,
			wantResult: map[string]float64{
				"metric1": 100,
				"metric2": 102,
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
				},
				storage: &storage.MemoryStorage{
					Metrics: map[string]float64{
						"metric2": 100,
					},
				},
			},
			wantCount: 1,
			wantResult: map[string]float64{
				"metric2": 102,
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
					Metrics: map[string]float64{
						"metric3": 100,
					},
				},
			},
			wantCount: 1,
			wantResult: map[string]float64{
				"metric1": 100,
			},
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
			if res.StatusCode == http.StatusOK {
				assert.Equal(t, len(tt.args.storage.IndexMetrics()), tt.wantCount)
				assert.True(t, reflect.DeepEqual(tt.wantResult, tt.args.storage.IndexMetrics()))
			}

		})
	}
}
