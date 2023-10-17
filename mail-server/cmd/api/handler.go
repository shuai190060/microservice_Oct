package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessge struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}
	var requestPayload mailMessge
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}
	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println(err)
		app.errorJson(w, err)
		return
	}
	payload := JsonResponse{
		Error:   false,
		Message: "sent to " + requestPayload.To,
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}
