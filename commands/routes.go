package main

import (
	"net/http"
)

//MuxHandler : mux handler function type definition
type MuxHandler func(w http.ResponseWriter, r *http.Request)

//Route : a type definition for route
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  MuxHandler
}

//Routes : All routes are defined here
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CreateCustomer",
		"POST",
		"/customer",
		CreateCustomer,
	},
	Route{
		"FindCustomer",
		"GET",
		"/customer/{email}",
		FindCustomer,
	},
	Route{
		"NotifyCustomer",
		"GET",
		"/notify/{email}",
		NotifyCustomer,
	},
}
