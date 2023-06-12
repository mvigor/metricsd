package apiclient

import (
	"fmt"
	"net/http"
)

type ApiClientInterface interface {
	SetServer(string)
	PostMetric(string, interface{})
}

type ApiHttpClient struct {
	server string
	client http.Client
}

func NewApiHttpClient(server string) *ApiHttpClient {
	c := new(ApiHttpClient)
	c.SetServer(server)
	return c
}

func (c *ApiHttpClient) SetServer(server string) {
	c.server = server

}

func (c *ApiHttpClient) PostMetric(sname string, value interface{}) {

	url := "http://" + c.server + "/update/" + sname + "/" + serializeData(value)
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
