package main

import (
	"github.com/mvigor/metricsd/internal/interfaces"
	router2 "github.com/mvigor/metricsd/internal/router"
	"log"
	"net/http"
)

func loadRoutingMap() interfaces.RoutingMap {
	endpoints := router2.Map
	return interfaces.RoutingMap{Endpoints: endpoints}
}

func main() {

	r := router2.ChiRouter{}
	mux, err := r.LoadRoutingTable(loadRoutingMap())

	if err != nil {
		panic(err)
	}

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
