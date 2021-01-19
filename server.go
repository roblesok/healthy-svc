package main

import (
	"log"
	"net/http"

	"github.com/justinas/alice"
	"github.com/roblesok/healthy-svc/handler"
	"github.com/roblesok/healthy-svc/middleware"
)

func main() {
	port := ":8080"

	commonHandlers := alice.New(middleware.Logger)

	http.Handle("/health", commonHandlers.ThenFunc(handler.Health))

	log.Fatal(http.ListenAndServe(port, nil))
}
