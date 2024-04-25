package main

import (
	"log"
	"prathameshj.dev/passhash/db"
	"prathameshj.dev/passhash/server"
)

func main() {
	db, err := db.NewDataBaseClient()
	if err != nil {
		log.Fatalf("failed to initialize db client: %s", err)
	}

	srv := server.StartServer(db)
	srv.Start()
}
