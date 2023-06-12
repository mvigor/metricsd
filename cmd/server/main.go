package main

import (
	"github.com/mvigor/metricsd/cmd/server/interfaces"
	"github.com/mvigor/metricsd/cmd/server/router"
	"log"
	"net/http"
)

func loadRoutingMap() interfaces.RoutingMap {
	endpoints := router.Map
	return interfaces.RoutingMap{Endpoints: endpoints}
}

func main() {

	r := router.ChiRouter{}
	mux, err := r.LoadRoutingTable(loadRoutingMap())

	if err != nil {
		panic(err)
	}

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
