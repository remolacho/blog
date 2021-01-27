package main

import (
	"blog/internal/server"
	"log"
)

// se crea desde aca
// https://dev.to/orlmonteverde/api-rest-con-go-golang-y-postgresql-m0o
func main() {
	serv, err := server.New()
	if err != nil { log.Fatal(err) }

	// start the server.
	go serv.Start()

	// Attempt a graceful shutdown press ctrl c.
	serv.Close()
}
