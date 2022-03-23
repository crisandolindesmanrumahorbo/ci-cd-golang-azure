package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
