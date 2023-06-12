package router

import (
	"github.com/mvigor/metricsd/internal/handlers"
	"github.com/mvigor/metricsd/internal/interfaces"
	"net/http"
)

var Map = []interfaces.RoutingEndpoint{
	{
		Method:      http.MethodGet,
		Pattern:     "/",
		Handler:     handlers.IndexHandler,
		Middlewares: interfaces.Middlewares{},
	},
	{
		Method:      http.MethodGet,
		Pattern:     "/{metric_name}",
		Handler:     handlers.ShowHandler,
		Middlewares: interfaces.Middlewares{},
	},
	{
		Method:      http.MethodPost,
		Pattern:     "/update/{metric_name}/{metric_value}",
		Handler:     handlers.UpdateHandler,
		Middlewares: interfaces.Middlewares{},
	},
}
