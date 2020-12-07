package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Datatypes
type Post struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

// Static Post data
var posts []Post
var eventbus = "http://localhost:8085/events"

func getNextId() int64 {
	return int64(len(posts))
}

func addPost(w http.ResponseWriter, r *http.Request) {
	// Decode
	decoder := json.NewDecoder(r.Body)
	var msg Message
	err := decoder.Decode(&msg)
	if err != nil {
		panic(err)
	}

	// Create Post
	var post Post
	post.ID = getNextId()
	post.Message = msg.Message
	posts = append(posts, post)

	log.Println(posts)

	// Create Response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		panic(err)
	}

	// Call eventbus
	event_req_body, err := json.Marshal(post)
	http.Post(eventbus, "application/json", bytes.NewReader(event_req_body))
}

// Handler to get all customers encoded as JSON
func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if err := json.NewEncoder(w).Encode(posts); err != nil {
			panic(err)
		}
	case "POST":
		addPost(w, r)
	}
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var msg Message
	decoder.Decode(&msg)
	log.Println("Received event")
	log.Println(msg)
}
