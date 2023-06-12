package router

import (
	handlers2 "github.com/mvigor/metricsd/cmd/server/handlers"
	"github.com/mvigor/metricsd/cmd/server/interfaces"
	"net/http"
)

var Map = []interfaces.RoutingEndpoint{
	{
		Method:      http.MethodGet,
		Pattern:     "/",
		Handler:     handlers2.IndexHandler,
		Middlewares: interfaces.Middlewares{},
	},
	{
		Method:      http.MethodGet,
		Pattern:     "/{metric_name}",
		Handler:     handlers2.ShowHandler,
		Middlewares: interfaces.Middlewares{},
	},
	{
		Method:      http.MethodPost,
		Pattern:     "/update/{metric_name}/{metric_value}",
		Handler:     handlers2.UpdateHandler,
		Middlewares: interfaces.Middlewares{},
	},
}
