package handlers

import (
	"net/http"

	"github.com/mvigor/metricsd/internal/entities"
	"github.com/mvigor/metricsd/internal/storage"
)

func ShowHandler(params map[string]string, storage storage.Storage) http.HandlerFunc {

	metricName := params["metric_name"]
	metricType := params["metric_type"]
	value, ok := storage.GetMetric(metricType, metricName)

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
		resp.Write([]byte(value.ToString()))
	}
}

func UpdateHandler(params map[string]string, storage storage.Storage) http.HandlerFunc {

	metricName := params["metric_name"]
	metricValue := params["metric_value"]
	metricType := params["metric_type"]

	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		metric, err := entities.MetricFactory(metricType, metricName, metricValue)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		err = storage.SetMetric(metric)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		resp.Header().Set("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("updated"))
	}
}
