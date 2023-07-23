package handlers

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/mvigor/metricsd/internal/storage"
)

func IndexHandler(params map[string]string, storage storage.Storage) http.HandlerFunc {
	return func(resp http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		metrics := storage.IndexMetrics()

		keys := make([]string, 0, len(metrics))
		for k := range storage.IndexMetrics() {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		resp.Header().Set("Content-Type", "text/html")
		for _, k := range keys {
			val := metrics[k]
			resp.Write([]byte(fmt.Sprintf("%s = %s<br>\n", val.GetName(), val.ToString())))
		}
		resp.WriteHeader(http.StatusOK)
	}
}
