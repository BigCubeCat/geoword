package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wordgeocode/coder"
)

func encodeCoordinates(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	address := r.URL.Query().Get("address")

	if len(latitude) == 0 || len(longitude) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Invalid latitude or longitude")); err != nil {
			fmt.Println(err)
		}
		return
	}
	if len(address) > 0 {
		// add address to database
	}

	resp := make(map[string]string)
	resp["result"] = coder.EncodeCoords(latitude, longitude)

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	if _, err := w.Write(jsonResp); err != nil {
		log.Println("error: ", err)
		return
	}
}
