package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//FindCustomerFromDB : given an email, it will find the customer in redis
func FindCustomerFromDB(email string) (Customer, error) {
	var customer Customer
	data, err := FindRecord("Customer:" + email)
	if err != nil {
		return customer, err
	}
	json.Unmarshal([]byte(data), &customer)
	return customer, nil
}

//FindCustomer : find a customer by email (in request)
func FindCustomer(w http.ResponseWriter, r *http.Request) {

	email := mux.Vars(r)["email"]
	fmt.Println("email", email)
	customer, err := FindCustomerFromDB(email)
	if err != nil {
		HandleError(err)
		fmt.Fprintf(w, "error ocurred")
		return
	}
	json.NewEncoder(w).Encode(customer)
}
