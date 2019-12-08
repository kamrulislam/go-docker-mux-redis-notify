package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Notifier : a generic Notifier (implementation independent)
type Notifier interface {
	Notify(customer Customer) error
}

//NotifyCustomer : NotifyCustomer
func NotifyCustomer(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	fmt.Println("email", email)

	customer, err := FindCustomerFromDB(email)
	if err != nil {
		HandleError(err)
		fmt.Fprintf(w, "error ocurred")
		return
	}
	// json.NewEncoder(w).Encode(customer)
	var notifier Notifier
	notifier = SendGridNotifier{}
	err = notifier.Notify(customer)
	if err != nil {
		HandleError(err)
		fmt.Fprintf(w, "error ocurred")
		return
	}
	jsonBody, _ := json.Marshal(map[string]bool{
		"success": true,
	})
	w.Write(jsonBody)
}
