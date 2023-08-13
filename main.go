package main

import (
	"wordgeocode/db"
	"wordgeocode/server"
)

func main() {
	ldb := db.LoadDataBase("data.ldb")
	serv := server.ApiServer{Port: "3333", Ldb: ldb}
	serv.Run()
}
