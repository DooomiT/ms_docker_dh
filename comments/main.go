package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/posts/comments", handleComment)
	router.HandleFunc("/events", handleEvent)
	log.Println("CustomerServer: Listening on http://localhost:8081/posts/comments ...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
