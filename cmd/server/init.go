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

func InitApp(addr string) error {
	r := router2.ChiRouter{}
	mux, err := r.LoadRoutingTable(loadRoutingMap())
	if err != nil {
		return err
	}
	log.Fatalln(http.ListenAndServe(addr, mux))

	return nil
}
