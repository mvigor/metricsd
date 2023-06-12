package handlers

import (
	"github.com/mvigor/metricsd/cmd/server/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowHandler(t *testing.T) {
	storage := storage.MemoryStorage{
		Metrics: map[string]string{
			"metric1": "1",
			"metric2": "value",
			"metric3": "value2",
			"metric4": "",
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
			expectedBody:        "metric name = metric1, value = 1",
			expectedContentType: "text/html",
		},
		{
			name:                "test case #2",
			method:              http.MethodGet,
			request:             map[string]string{"metric_name": "metric3"},
			expectedCode:        http.StatusOK,
			expectedBody:        "metric name = metric3, value = value2",
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
				resBody, err := io.ReadAll(res.Body)

				require.NoError(t, err)

				assert.Equal(t, tt.expectedBody, string(resBody))
				assert.Equal(t, tt.expectedContentType, res.Header.Get("Content-Type"))
			}
		})

	}
}
