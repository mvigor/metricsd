package apiclient

import (
	"fmt"

	"github.com/mvigor/metricsd/internal/entities"

	"net/http"
)

type APIClient interface {
	PostMetric(string, entities.MetricValue) error
}

type HttpApiClient struct {
	server string
	client http.Client
}

func NewHttpApiClient(server string) APIClient {
	return &HttpApiClient{
		server: server,
	}
}

func (c *HttpApiClient) PostMetric(sname string, value entities.MetricValue) error {

	url := fmt.Sprintf("http://%s/update/%s/%s/%s", c.server, value.Type, sname, serializeData(value.Value))
	resp, err := c.client.Post(url, "text/html", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
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
