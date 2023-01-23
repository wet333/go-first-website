package main

import (
	"fmt"
	kvStore "go-first-website/kvStore"
	"net/http"
)

var kvStoreInstance = kvStore.KeyValueStore{}

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
	for id, message := range kvStoreInstance.GetList() {
		_, _ = fmt.Fprintf(w, "ID: %s, Message: %s\n", id, message)
	}
}

func saveMessage(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	id := r.FormValue("id")

	kvStoreInstance.Add(id, message)
	saveMessages()
}

func handleSingleMessage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idWithComillas := "\"" + id + "\""

	message, ok := kvStoreInstance.Get(idWithComillas)
	if !ok {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "ID: %s, Message: %s", id, message)
}

func loadMessages() {
	kvStoreInstance.Load("messages.txt")
}

func saveMessages() {
	kvStoreInstance.Save("messages.txt")
}
