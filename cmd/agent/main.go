package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/mvigor/metricsd/cmd/agent/apiclient"
	"github.com/mvigor/metricsd/cmd/agent/collectors"
	"github.com/mvigor/metricsd/internal/utils"
	"time"
)

type Config struct {
	Address        string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PoolInterval   int    `env:"POLL_INTERVAL"`
}

const DefaultServer = "localhost:8080"

func main() {

	server, poolInterval, reportInterval := getParameters()

	client := apiclient.NewAPIHttpClient(server)
	collectors.StartCollectors(time.Duration(poolInterval) * time.Second)
	for {
		time.Sleep(time.Duration(reportInterval) * time.Second)
		res := collectors.CollectData()

		for key, value := range res {
			fmt.Printf("%s = %v\n", key, value)
			client.PostMetric(key, value)
		}
		fmt.Println("--------------------------------------------------")
	}
}

func getParameters() (endPointServer string, poolInterval int, reportInterval int) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	endPointServer = DefaultServer

	addr := new(utils.NetAddress)
	_ = flag.Value(addr)
	flag.Var(addr, "a", "Net address host:port")
	flag.IntVar(&reportInterval, "r", 10, "report period interval")
	flag.IntVar(&poolInterval, "p", 2, "pool interval")

	flag.Parse()

	if len(addr.String()) > 2 {
		endPointServer = addr.String()
	}
	if len(cfg.Address) > 0 {
		endPointServer = cfg.Address
	}

	if cfg.ReportInterval > 0 {
		reportInterval = cfg.ReportInterval
	}

	if cfg.PoolInterval > 0 {
		poolInterval = cfg.PoolInterval
	}

	return
}
