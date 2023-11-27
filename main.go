package main

import (
	"database/sql"
	"log"

	"github.com/Justinecav/simplebank.git/api"
	db "github.com/Justinecav/simplebank.git/db/sqlc"
	"github.com/Justinecav/simplebank.git/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("error in loading config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the database :", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: %w ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server :", err)
	}
}
