package main

import (
	"log"
	"net/http"

	"github.com/roblesok/healthy/handler"
)

func main() {
	port := ":8080"

	http.HandleFunc("/health", handler.Health)

	log.Fatal(http.ListenAndServe(port, nil))
}
