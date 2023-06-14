package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/mvigor/metricsd/internal/interfaces"
	router2 "github.com/mvigor/metricsd/internal/router"
	"github.com/mvigor/metricsd/internal/utils"
	"log"
	"net/http"
)

const DefaultServer = "localhost:8080"

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

func loadRoutingMap() interfaces.RoutingMap {
	endpoints := router2.Map
	return interfaces.RoutingMap{Endpoints: endpoints}
}

func InitApp(addr string) error {
	r := router2.ChiRouter{}
	mux, err := r.LoadRoutingTable(loadRoutingMap())
	if err != nil {
		return err
	}
	log.Fatalln(http.ListenAndServe(addr, mux))

	return nil
}
