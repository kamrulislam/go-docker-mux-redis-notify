package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

//SendGridNotifier : a helper struct to implement Notifier interface
type SendGridNotifier struct {
}

//Notify : is a receiver function which also implements Notifier interface
func (s SendGridNotifier) Notify(customer Customer) error {
	//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
	apiKey := os.Getenv("EMAIL_API_KEY")
	fmt.Println("EMAIL_API_KEY", apiKey)
	// ideally the API key would be coming from environment
	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = []byte(` {
		"personalizations": [
			{
				"to": [
					{
						"email": "` + customer.Email + `"
					}
				],
				"subject": "Sending notification with Twilio SendGrid"
			}
		],
		"from": {
			"email": "varun.verma@bcgdv.com"
		},
		"content": [
			{
				"type": "text/plain",
				"value": "This is to notify you that BCGDV is testing email sending."
			}
		]
	}`)
	_, err := sendgrid.API(request)

	return err
}
