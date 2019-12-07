package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CreateCustomer : add a customer to redis
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	Logger(r)

	var newCustomer Customer
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the customer")
		return
	}

	json.Unmarshal(reqBody, &newCustomer)

	data, err := json.Marshal(newCustomer)
	if err != nil {
		fmt.Println(err)
	}

	err = AddRecord("Customer:"+newCustomer.Email, data)
	if err != nil {
		HandleError(err)
		fmt.Fprint(w, "Failed to process")
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newCustomer); err != nil {
		HandleError(err)
	}
}
