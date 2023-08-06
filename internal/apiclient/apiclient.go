package apiclient

import (
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

	url := fmt.Sprintf("http://%s/update/%s/%s/%s", c.server, value.Type, sname, serializeData(value.Value))
	resp, err := c.client.Post(url, "text/html", nil)
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

func serializeData(value interface{}) string {
	switch value.(type) {
	case float64, float32:
		return fmt.Sprintf("%g", value)
	default:
		return fmt.Sprintf("%d", value)
	}
}
