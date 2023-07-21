package router

import (
	"net/http"

	"github.com/mvigor/metricsd/internal/handlers"
	"github.com/mvigor/metricsd/internal/interfaces"
)

var Map = []interfaces.RoutingEndpoint{
	{
		Method:  http.MethodGet,
		Pattern: "/",
		Handler: handlers.IndexHandler,
	},
	{
		Method:  http.MethodGet,
		Pattern: "/value/{metric_type}/{metric_name}",
		Handler: handlers.ShowHandler,
	},
	{
		Method:  http.MethodPost,
		Pattern: "/update/{metric_type}/{metric_name}/{metric_value}",
		Handler: handlers.UpdateHandler,
	},
}
