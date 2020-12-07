package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Datatypes
type Comment struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

// Static comment data
var comments_by_id = make(map[int64][]Comment)

var eventbus = "http://localhost:8085/events"

func getNextId(post_id int64) int64 {
	if comments_by_id[post_id] == nil {
		return 0
	}
	return int64(len(comments_by_id[post_id]))
}

func addComment(w http.ResponseWriter, r *http.Request) {
	// Decode
	decoder := json.NewDecoder(r.Body)
	var msg Message
	err := decoder.Decode(&msg)
	if err != nil {
		panic(err)
	}
	post_id := r.URL.Query()["id"]
	post_id_64, _ := strconv.ParseInt(post_id[0], 10, 64)
	log.Println(post_id_64)

	// Create Comment
	var comment Comment
	comment.ID = getNextId(post_id_64)
	comment.Message = msg.Message
	log.Println(comment)
	if comments_by_id[post_id_64] == nil {
		comments_by_id[post_id_64] = []Comment{comment}
	} else {
		comments_by_id[post_id_64] = append(comments_by_id[post_id_64], comment)
	}

	log.Println(comments_by_id)

	// Create Response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comments_by_id)
	if err != nil {
		panic(err)
	}

	// Call eventbus
	event_req_body, err := json.Marshal(comment)
	http.Post(eventbus, "application/json", bytes.NewReader(event_req_body))
}

// Handler to get all customers encoded as JSON
func handleComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if err := json.NewEncoder(w).Encode(comments_by_id); err != nil {
			panic(err)
		}
	case "POST":
		addComment(w, r)
	}
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var msg Message
	decoder.Decode(&msg)
	log.Println("Received event")
	log.Println(msg)
}
