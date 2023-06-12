package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/mvigor/metricsd/internal/utils"
)

const DefaultServer = "localhot:8080"

type Config struct {
	Address string `env:"ADDRESS"`
}

func main() {

	var serverEndpoint = DefaultServer
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	addr := new(utils.NetAddress)
	_ = flag.Value(addr)
	flag.Var(addr, "a", "Net address host:port")
	flag.Parse()

	if len(addr.String()) > 2 {
		serverEndpoint = addr.String()
	}

	if len(cfg.Address) > 2 {
		serverEndpoint = cfg.Address
	}

	err = InitApp(serverEndpoint)
	if err != nil {
		panic(fmt.Sprintf("couldn't start application\n%s", err.Error()))
	}

}
