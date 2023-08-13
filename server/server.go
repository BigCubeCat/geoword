package server

import (
	"log"
	"net/http"
	"wordgeocode/db"
)

type ApiServer struct {
	Ldb  *db.LevelDbDriver
	Port string
}

func (server *ApiServer) EncodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		word := encodeCoordinates(r)
		makeEncodeResponse(w, WordDTO{Word: word})
	case "POST":
		word := encodeCoordinates(r)
		address := getAddress(r)
		if err := server.Ldb.Set(word, address); err != nil {
			log.Println("cant set")
			log.Println(err)
		}
		makeEncodeResponse(w, WordDTO{Word: word})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *ApiServer) DecodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		dto := AddressDTO{}
		word := getWord(r)
		dto.Latitude, dto.Longitude = decodeCoordinates(word)
		dto.Address = server.Ldb.Get(word)
		makeDecodeResponse(w, dto)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *ApiServer) Run() {
	http.HandleFunc("/encode", server.EncodeHandler)
	http.HandleFunc("/decode", server.DecodeHandler)

	if err := http.ListenAndServe(":"+server.Port, nil); err != nil {
		log.Fatal(err)
	}
}
