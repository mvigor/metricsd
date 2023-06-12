package handlers

import (
	"fmt"
	"github.com/mvigor/metricsd/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	storage := storage.MemoryStorage{
		Metrics: map[string]string{
			"metric1": "1",
			"metric2": "value",
			"metric3": "value2",
			"metric4": "",
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
		successBody += fmt.Sprintf("%s = %s<br>\n", k, val)
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
				resBody, err := io.ReadAll(res.Body)

				require.NoError(t, err)

				assert.Equal(t, tt.expectedBody, string(resBody))
				assert.Equal(t, tt.expectedContentType, res.Header.Get("Content-Type"))
			}
		})

	}
}
