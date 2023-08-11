package server

import (
	"fmt"
	"io"
	"net/http"
)

func decodeCoordinates(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	address := r.URL.Query().Get("address")
	if len(address) > 0 {
		// Search address
		fmt.Println("search...")
	}
	if _, err := io.WriteString(w, word+"\n"); err != nil {
		fmt.Println(err)
	}
}
