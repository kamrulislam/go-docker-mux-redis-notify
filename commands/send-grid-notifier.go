package main

import (
	"github.com/sendgrid/sendgrid-go"
)

type SendGridNotifier struct {
}

//Notify : is a receiver function which also implements Notifier interface
func (s SendGridNotifier) Notify(customer Customer) error {
	request := sendgrid.GetRequest("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", "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = []byte(` {
		"personalizations": [
			{
				"to": [
					{
						"email": "` + customer.Email + `"
					}
				],
				"subject": "Sending with Twilio SendGrid is Fun"
			}
		],
		"from": {
			"email": "kamruli@gmail.com"
		},
		"content": [
			{
				"type": "text/plain",
				"value": "and easy to do anywhere, even with Go"
			}
		]
	}`)
	_, err := sendgrid.API(request)

	return err
}
