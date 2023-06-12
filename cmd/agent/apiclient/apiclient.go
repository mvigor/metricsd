package apiclient

import "net/http"

type ApiClientInterface interface {
	SetServer(string)
	PostMetric(string, string)
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

func (c *ApiHttpClient) PostMetric(sname, stype string) {

}
