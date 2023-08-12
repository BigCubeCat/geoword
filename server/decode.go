package server

import (
	"encoding/json"
	"log"
	"net/http"
	"wordgeocode/coder"
)

func decodeCoordinates(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	word := r.URL.Query().Get("word")

	resp["result"] = coder.DecodeCoords(word)
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	if _, err := w.Write(jsonResp); err != nil {
		log.Println("error: ", err)
		return
	}
}
