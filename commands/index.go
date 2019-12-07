package main

import (
	"fmt"
	"net/http"
)

//Index : an entry api
func Index(w http.ResponseWriter, r *http.Request) {

	Logger(r)

	fmt.Fprintf(w, "Welcome to Kamrul's API server")
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//fmt.Fprintf(w, "Hello, %s\n", p.ByName("anything"))
}
