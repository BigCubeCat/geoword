package main

import (
	"log"
	"net/http"
	"wordgeocode/coder"
	"wordgeocode/server"
)

func main() {
	coder.Prepare()
	http.HandleFunc("/encode", server.EncodeHandler)
	http.HandleFunc("/decode", server.DecodeHandler)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}
