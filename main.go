package main

import (
	"log"
	"net/http"
	"wordgeocode/server"
)

func main() {
	http.HandleFunc("/encode", server.EncodeHandler)
	http.HandleFunc("/decode", server.DecodeHandler)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}
