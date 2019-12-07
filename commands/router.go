package main

import "github.com/gorilla/mux"

//CreateRouter : router is initialized here
func CreateRouter() *mux.Router {

	router := mux.NewRouter()

	for _, route := range routes {

		router.HandleFunc(route.Pattern, route.Handle).Methods(route.Method).Name(route.Name)
	}

	return router
}
