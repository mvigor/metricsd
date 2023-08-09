package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/mvigor/metricsd/internal/entities"
	"github.com/mvigor/metricsd/internal/utils"

	"net/http"
)

type APIClient interface {
	PostMetric(string, entities.MetricValue) error
}

type HTTPAPIClient struct {
	server string
	client http.Client
}

func NewHTTPAPIClient(server string) APIClient {
	return &HTTPAPIClient{
		server: server,
	}
}

func (c *HTTPAPIClient) PostMetric(sname string, value entities.MetricValue) error {

	metric := entities.ApiMetrics{
		ID:    sname,
		MType: string(value.Type),
	}

	switch value.Type {
	case entities.GAUGE:
		val := value.Value.(float64)
		metric.Value = &val
		break
	case entities.COUNTER:
		val := value.Value.(int64)
		metric.Delta = &val
		break
	}

	url := fmt.Sprintf("http://%s/update/", c.server)

	jsonBody, _ := json.Marshal(metric)

	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	logger := utils.GetLogger()
	sugar := logger.Sugar()
	defer sugar.Desugar()

	sugar.Infoln(
		"code", resp.StatusCode,
		"size", resp.ContentLength,
	)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned status code %d", resp.StatusCode)
	}

	return nil
}
