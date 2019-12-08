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

//HandleError : a generic function to handle error
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
