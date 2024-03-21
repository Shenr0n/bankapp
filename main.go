package main

import (
	"database/sql"
	"log"

	"github.com/Shenr0n/bankapp/api"
	db "github.com/Shenr0n/bankapp/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	// Establish connection to the db
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
