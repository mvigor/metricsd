package main

import (
	"fmt"
	"github.com/mvigor/metricsd/cmd/agent/apiclient"
	"github.com/mvigor/metricsd/cmd/agent/collectors"
	"time"
)

func main() {
	client := apiclient.NewApiHttpClient("localhost:8080")
	for {
		res := collectors.CollectData()

		for key, value := range res {
			fmt.Printf("%s = %v\n", key, value)
			client.PostMetric(key, value)
		}
		fmt.Println("--------------------------------------------------")

		time.Sleep(5 * time.Second)
	}
}
