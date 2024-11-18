package main

import (
	"log"
	"luizalabs-chalenge/infra/server"
)

func main() {
	server := server.NewServer()
	log.Fatalf("fatal error while running app server (%v)", server.Run())
}
