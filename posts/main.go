package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/posts", handlePost)
	router.HandleFunc("/events", handleEvent)
	log.Println("CustomerServer: Listening on http://localhost:8080/posts ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
