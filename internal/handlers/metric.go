package handlers

import (
	"fmt"
	"github.com/mvigor/metricsd/internal/storage"
	"net/http"
)

func ShowHandler(params map[string]string, storage storage.Storage) http.HandlerFunc {

	metricName := params["metric_name"]
	value, ok := storage.GetMetric(metricName)

	return func(resp http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		resp.Header().Set("Content-Type", "text/html")
		if !ok {
			resp.WriteHeader(http.StatusNotFound)
			return
		}

		resp.WriteHeader(http.StatusOK)
		r := fmt.Sprintf("metric name = %s, value = %v", metricName, value)
		resp.Write([]byte(r))
	}
}

func UpdateHandler(params map[string]string, storage storage.Storage) http.HandlerFunc {

	metricName := params["metric_name"]
	metricValue := params["metric_value"]

	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := storage.SetMetric(metricName, metricValue)
		resp.Header().Set("Content-Type", "text/html")
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.WriteHeader(http.StatusOK)
	}
}
