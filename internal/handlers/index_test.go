package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/mvigor/metricsd/internal/entities"
	"github.com/mvigor/metricsd/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexHandler(t *testing.T) {
	storage := storage.MemoryStorage{
		Metrics: map[string]entities.Metric{
			"counter_metric1": entities.NewMetricCounter("metric1", 11),
			"counter_metric2": entities.NewMetricCounter("metric2", 100),
			"counter_metric3": entities.NewMetricCounter("metric3", 200),
			"counter_metric4": entities.NewMetricCounter("metric4", 202),
		},
	}

	successBody := ""
	keys := make([]string, 0, len(storage.Metrics))
	for k := range storage.Metrics {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		val := storage.Metrics[k]
		successBody += fmt.Sprintf("%s = %s<br>\n", val.GetName(), val.ToString())
	}

	testCases := []struct {
		name                string
		method              string
		expectedCode        int
		expectedBody        string
		expectedContentType string
	}{
		{
			name:                "test case #1",
			method:              http.MethodGet,
			expectedCode:        http.StatusOK,
			expectedBody:        successBody,
			expectedContentType: "text/html",
		},
		{
			name:         "test case #2",
			method:       http.MethodPut,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "test case #3",
			method:       http.MethodDelete,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "test case #4",
			method:       http.MethodPost,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", nil)
			w := httptest.NewRecorder()
			handlerFunc := IndexHandler(nil, &storage)
			handlerFunc(w, request)
			res := w.Result()
			assert.Equal(t, res.StatusCode, tt.expectedCode)

			if res.StatusCode == http.StatusOK {
				defer res.Body.Close()
				resBody, err := io.ReadAll(res.Body)
				defer require.NoError(t, err)

				assert.Equal(t, tt.expectedBody, string(resBody))
				assert.Equal(t, tt.expectedContentType, res.Header.Get("Content-Type"))
			}
		})

	}
}
