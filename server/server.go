package server

import (
	"net/http"
)

func EncodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		encodeCoordinates(w, r)
	case "POST":
		encodeCoordinates(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		decodeCoordinates(w, r)
	case "POST":
		decodeCoordinates(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
