package server

import (
	"fmt"
	"io"
	"net/http"
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
	res := latitude + "," + longitude
	if _, err := io.WriteString(w, res+"\n"); err != nil {
		fmt.Println(err)
	}
}
