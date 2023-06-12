package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/mvigor/metricsd/internal/interfaces"
	"github.com/mvigor/metricsd/internal/storage"
	"net/http"
)

var stor storage.Storage

type ChiRouter struct {
}

func (ChiR *ChiRouter) customHandler(action interfaces.RoutingEndpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := chi.RouteContext(r.Context()).URLParams
		paramsMap := map[string]string{}
		for index, key := range params.Keys {
			paramsMap[key] = params.Values[index]
		}

		action.Handler(paramsMap, stor)(w, r)
	}
}

func (ChiR *ChiRouter) LoadRoutingTable(table interfaces.RoutingMap) (http.Handler, error) {

	chiRouter := chi.NewRouter()

	if len(table.Endpoints) == 0 {
		return nil, fmt.Errorf("invalid routing table")
	}

	for _, action := range table.Endpoints {
		chiRouter.Route(action.Pattern, func(r chi.Router) {
			if len(action.Middlewares) > 0 {
				r.Use(action.Middlewares...)
			}
			r.Method(action.Method, "/", ChiR.customHandler(action))
		})
	}

	stor = storage.GetStorage()

	return chiRouter, nil
}

func (ChiR *ChiRouter) GetURLParam(req *http.Request, key string) string {
	return chi.URLParam(req, key)
}
