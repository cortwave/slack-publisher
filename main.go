package main

import (
	"encoding/json"
	"log"
	"net/http"
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
		url := "https://slack.com/api/chat.postMessage?token=" + os.Getenv("TOKEN") +
			"&channel=" + message.Channel + "&text=" + message.Text
		log.Println("request on", url)
		resp, err := http.Post(url, "text/plain", nil)
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
