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

	commonHandlers := alice.New(middleware.Logger, middleware.Recovery)

	http.Handle("/health", commonHandlers.ThenFunc(handler.Health))

	http.Handle("/fail", commonHandlers.ThenFunc(handler.Fail))

	log.Fatal(http.ListenAndServe(port, nil))
}
