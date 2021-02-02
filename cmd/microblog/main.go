package main

import (
	"blog/internal/entities/post"
	"blog/internal/entities/user"
	"blog/internal/server"
	"blog/pkg/engineDB"
	"log"
)

// se crea desde aca
// https://dev.to/orlmonteverde/api-rest-con-go-golang-y-postgresql-m0o
func main() {
	serv, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database.
	factory := engineDB.Factory()
	if err := factory.DB.Error; err != nil {
		log.Fatal(err)
	}

	factory.DB.AutoMigrate(&user.User{}, &post.Post{})

	// start the server.
	go serv.Start()

	// Attempt a graceful shutdown press ctrl c.
	serv.Close()
}
