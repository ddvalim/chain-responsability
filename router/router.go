package router

import (
	"chain-responsability/router/routes"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	return ConfigRouter(router)
}

func ConfigRouter(router *mux.Router) *mux.Router {
	r := routes.UserRoutes

	for _, route := range r {
		router.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return router
}
