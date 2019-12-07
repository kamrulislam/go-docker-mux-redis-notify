package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//FindCustomer : find a customer by email
func FindCustomer(w http.ResponseWriter, r *http.Request) {

	email := mux.Vars(r)["email"]
	fmt.Println("email", email)
	data, err := FindRecord("Customer:" + email)
	if err != nil {
		HandleError(err)
		fmt.Fprintf(w, "error ocurred")
		return
	}
	var customer Customer
	json.Unmarshal([]byte(data), &customer)
	json.NewEncoder(w).Encode(customer)
}
