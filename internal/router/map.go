package router

import (
	"net/http"

	"github.com/mvigor/metricsd/internal/handlers"
	"github.com/mvigor/metricsd/internal/interfaces"
	"github.com/mvigor/metricsd/internal/middlewares"
)

var Map = []interfaces.RoutingEndpoint{
	{
		Method:  http.MethodGet,
		Pattern: "/",
		Handler: handlers.IndexHandler,
		Middlewares: interfaces.Middlewares{
			middlewares.WithLogging,
		},
	},
	{
		Method:  http.MethodGet,
		Pattern: "/value/{metric_type}/{metric_name}",
		Handler: handlers.ShowHandler,
		Middlewares: interfaces.Middlewares{
			middlewares.WithLogging,
		},
	},
	{
		Method:  http.MethodPost,
		Pattern: "/update/{metric_type}/{metric_name}/{metric_value}",
		Handler: handlers.UpdateHandler,
		Middlewares: interfaces.Middlewares{
			middlewares.WithLogging,
		},
	},
}
