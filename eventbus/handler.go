package main

import (
	"fmt"
	"net/http"
)

var service_urls = [...]string{
	"http://localhost:8080/events",
	"http://localhost:8081/events"}

func forwardToServices(r *http.Request) {
	event := r.Body
	for _, url := range service_urls {
		http.Post(url, "application/json", event)
	}
}

// Handler to get all customers encoded as JSON
func handleEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "No such Method 'GET'")
	case "POST":
		forwardToServices(r)
	}
}
