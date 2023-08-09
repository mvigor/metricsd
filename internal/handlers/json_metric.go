package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mvigor/metricsd/internal/entities"
	"github.com/mvigor/metricsd/internal/storage"
)

func UpdateJsonHandler(params map[string]string, storage storage.Storage) http.HandlerFunc {

	return func(resp http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			resp.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var jsonMetric entities.ApiMetrics
		var metric entities.Metric

		jsonDecoder := json.NewDecoder(req.Body)
		err := jsonDecoder.Decode(&jsonMetric)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		metric, err = entities.MetricJsonFactory(jsonMetric)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		err = storage.SetMetric(metric)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("updated"))
	}
}
