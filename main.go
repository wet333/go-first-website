package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var messages = make(map[string]string)

func main() {
	loadMessages()

	http.HandleFunc("/", homePage)
	http.HandleFunc("/messages", handleMessages)
	http.HandleFunc("/get_message", handleSingleMessage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

	fmt.Println("Listening on port :8080")

	defer saveMessages()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to the message API!")
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showAllMessages(w, r)
	case "POST":
		saveMessage(w, r)
	}
}

func showAllMessages(w http.ResponseWriter, r *http.Request) {
	for id, message := range messages {
		_, _ = fmt.Fprintf(w, "ID: %s, Message: %s\n", id, message)
	}
}

func saveMessage(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	id := r.FormValue("id")

	messages[id] = message
	saveMessages()
}

func handleSingleMessage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idWithComillas := "\"" + id + "\""

	message, ok := messages[idWithComillas]
	if !ok {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "ID: %s, Message: %s", id, message)
}

func loadMessages() {
	data, err := ioutil.ReadFile("messages.txt")
	if err != nil {
		return
	}
	json.Unmarshal(data, &messages)
}

func saveMessages() {
	data, _ := json.Marshal(messages)
	ioutil.WriteFile("messages.txt", data, 0644)
}
