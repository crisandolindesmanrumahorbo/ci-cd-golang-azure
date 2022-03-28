package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}
