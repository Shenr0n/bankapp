package main

import (
	"database/sql"
	"log"

	"github.com/Shenr0n/bankapp/api"
	db "github.com/Shenr0n/bankapp/db/sqlc"
	"github.com/Shenr0n/bankapp/util"
	_ "github.com/lib/pq"
)

/*const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)*/

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	// Establish connection to the db
	//conn, err := sql.Open(dbDriver, dbSource)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}
	//err = server.Start(serverAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
