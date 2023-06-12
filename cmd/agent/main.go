package main

import (
	"fmt"
	"github.com/mvigor/metricsd/cmd/agent/collectors"
	"time"
)

func main() {
	for {
		res := collectors.CollectData()

		for key, value := range res {
			fmt.Printf("%s = %v\n", key, value)
		}
		fmt.Println("--------------------------------------------------")

		time.Sleep(5 * time.Second)
	}
}
