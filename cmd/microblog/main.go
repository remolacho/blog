package main

import (
	"blog/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	serv, err := server.New("8000")
	if err != nil {
		log.Fatal(err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	serv.Close()
}