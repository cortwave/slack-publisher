package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

var publisher = make(chan publishRequest)

type publishRequest struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func main() {
	log.Println("Slack publsher started")
	go startSender()
	startServer()
}

func startSender() {
	for {
		message := <-publisher
		log.Println("message accepted", message)
		slackURL := "https://slack.com/api/chat.postMessage?token=" + os.Getenv("TOKEN") +
			"&channel=" + url.QueryEscape(message.Channel) + "&text=" + url.QueryEscape(message.Text)
		log.Println("request on", slackURL)
		resp, err := http.Post(slackURL, "text/plain", nil)
		if err != nil {
			log.Println("error", err)
		} else {
			log.Println(resp)
		}
	}
}

func startServer() {
	http.HandleFunc("/publish", publish)
	http.ListenAndServe(":8000", nil)
}

func publish(w http.ResponseWriter, r *http.Request) {
	log.Println("Publish request")
	request := publishRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	publisher <- request
}
