package handlers

import (
	"fmt"
	"github.com/mvigor/metricsd/internal/storage"
	"net/http"
	"sort"
)

func IndexHandler(params map[string]string, storage storage.Storage) http.HandlerFunc {
	return func(resp http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		keys := make([]string, 0, len(storage.IndexMetrics()))
		for k := range storage.IndexMetrics() {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		resp.Header().Set("Content-Type", "text/html")
		for _, k := range keys {
			val, _ := storage.GetMetric(k)
			resp.Write([]byte(fmt.Sprintf("%s = %s<br>\n", k, val)))
		}
		resp.WriteHeader(http.StatusOK)
	}
}
