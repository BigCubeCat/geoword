package server

import (
	"encoding/json"
	"net/http"
	"wordgeocode/coder"
)

func decodeCoordinates(word string) (string, string) {
	return coder.DecodeCoords(word)
}

func getWord(r *http.Request) string {
	return r.URL.Query().Get("word")
}

func makeDecodeResponse(w http.ResponseWriter, dto AddressDTO) {
	w.Header().Set("Content-Type", "application/json")
	if len(dto.Latitude)*len(dto.Longitude) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
