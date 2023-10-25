package main

import (
	"go-base-template/db"
	"go-base-template/src/web/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbMongo := db.Init(os.Getenv("MYSQL_CONNECTION"))

	server.Run(dbMongo)
}
