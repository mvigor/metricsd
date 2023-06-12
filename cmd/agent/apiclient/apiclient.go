package apiclient

import (
	"fmt"
	"net/http"
)

type APIClientInterface interface {
	SetServer(string)
	PostMetric(string, interface{})
}

type APIHttpClient struct {
	server string
	client http.Client
}

func NewAPIHttpClient(server string) *APIHttpClient {
	c := new(APIHttpClient)
	c.SetServer(server)
	return c
}

func (c *APIHttpClient) SetServer(server string) {
	c.server = server

}

func (c *APIHttpClient) PostMetric(sname string, value interface{}) {

	url := fmt.Sprintf("http://%s/update/%s/%s", c.server, sname, serializeData(value))
	resp, err := c.client.Post(url, "text/html", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("something wrong")
	}
}

func serializeData(value interface{}) string {
	switch value.(type) {
	case float64, float32:
		return fmt.Sprintf("%g", value)
	default:
		return fmt.Sprintf("%d", value)
	}
}
