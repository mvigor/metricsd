package interfaces

import (
	"github.com/mvigor/metricsd/cmd/server/storage"
	"net/http"
)

type RoutingMap struct {
	Endpoints []RoutingEndpoint
}

type Middlewares []func(http.Handler) http.Handler
type Handler func(map[string]string, storage.Storage) http.HandlerFunc

type RoutingEndpoint struct {
	Method      string
	Pattern     string
	Handler     Handler
	Middlewares Middlewares
}

type Router interface {
	LoadRoutingTable(table RoutingMap) (http.Handler, error)
}
