package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/mvigor/metricsd/internal/apiclient"
	"github.com/mvigor/metricsd/internal/collectors"
)

type Config struct {
	Address        string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PollInterval   int    `env:"POLL_INTERVAL"`
}

const DefaultServer = "localhost:8080"

func main() {

	cfg := getParameters()

	client := apiclient.NewHttpApiClient(cfg.Address)
	collectors.StartCollectors(time.Duration(cfg.PollInterval) * time.Second)
	for {
		time.Sleep(time.Duration(cfg.ReportInterval) * time.Second)
		res := collectors.CollectData()

		for key, value := range res {
			fmt.Printf("%s = %v\n", key, value)
			client.PostMetric(key, value)
		}
	}
}

func getParameters() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	if len(cfg.Address) < 2 {
		flag.StringVar(&cfg.Address, "a", DefaultServer, "Net address host:port")
	}

	flag.IntVar(&cfg.ReportInterval, "r", 10, "report period interval")
	flag.IntVar(&cfg.PollInterval, "p", 2, "poll interval")

	flag.Parse()

	return cfg
}
