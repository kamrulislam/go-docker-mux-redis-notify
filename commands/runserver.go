package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := CreateRouter()

	fmt.Println("Server listening!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
