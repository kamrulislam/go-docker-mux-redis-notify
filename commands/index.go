package main

import (
	"fmt"
	"net/http"
)

//Index : an entry api
func Index(w http.ResponseWriter, r *http.Request) {

	Logger(r)

	fmt.Fprintf(w, "Welcome to Kamrul's API server")
}
