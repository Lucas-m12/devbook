package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

func Config(router *mux.Router) *mux.Router {
	routes := usersRoutes
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return router
}
