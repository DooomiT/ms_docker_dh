package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/events", handleEvent)
	log.Println("CustomerServer: Listening on http://localhost:8085/events ...")
	log.Fatal(http.ListenAndServe(":8085", router))
}
