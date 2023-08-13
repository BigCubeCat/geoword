package server

import (
	"encoding/json"
	"net/http"
	"wordgeocode/coder"
)

func encodeCoordinates(r *http.Request) string {
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	if len(latitude) == 0 || len(longitude) == 0 {
		return ""
	}
	return coder.EncodeCoords(latitude, longitude)
}

func getAddress(r *http.Request) string {
	return r.URL.Query().Get("address")
}

func makeEncodeResponse(w http.ResponseWriter, dto WordDTO) {
	w.Header().Set("Content-Type", "application/json")
	if len(dto.Word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
