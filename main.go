package main

import (
	"database/sql"
	"log"

	"github.com/Justinecav/simplebank.git/api"
	db "github.com/Justinecav/simplebank.git/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:1926@localhost:5432/simple_bank?sslmode=disable"
	address  = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to the database :", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("Cannot start server :", err)
	}
}
