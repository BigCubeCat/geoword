package main

import (
	"log"
	"os"
	"wordgeocode/db"
	"wordgeocode/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbPath := os.Getenv("DATABASE")

	serv := server.ApiServer{
		Port: port,
		Ldb:  &db.LevelDbDriver{Path: dbPath},
	}
	serv.Ldb.Init()
	serv.Run()
}
