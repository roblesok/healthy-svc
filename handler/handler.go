package handler

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(map[string]interface{}{"status": "healthy"})

}

func Fail(_ http.ResponseWriter, _ *http.Request) {
	panic("Something went Wrong")
}
