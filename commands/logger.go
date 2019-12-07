package main

import (
	"log"
	"net/http"
	"time"
)

//Logger : a generic log function
func Logger(r *http.Request) {

	start := time.Now()

	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		//name,
		time.Since(start),
	)
}
