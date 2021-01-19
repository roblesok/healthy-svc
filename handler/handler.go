package handler

import "net/http"

func Health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello!\b"))
}

func Fail(_ http.ResponseWriter, _ *http.Request) {
	panic("Something went Wrong")
}
