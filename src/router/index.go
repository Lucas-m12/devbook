package router

import (
	"devbook/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	router := mux.NewRouter()
	return routes.Config(router)
}
