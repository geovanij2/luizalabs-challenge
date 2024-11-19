package main

import (
	"log"
	"luizalabs-chalenge/infra/database"
	"luizalabs-chalenge/infra/server"
)

func main() {
	database, err := database.Connect()
	if err != nil {
		log.Fatalf("error while connecting to database (%v)", err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalf("error while pinging database (%v)", err)
	}
	defer database.Close()

	server := server.NewServer()
	log.Fatalf("fatal error while running app server (%v)", server.Run())
}
