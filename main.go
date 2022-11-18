package main

import (
	"database/sql"
	"log"

	"github.com/epivoca/balance_service/api"
	db "github.com/epivoca/balance_service/db/sqlc"
	"github.com/epivoca/balance_service/utils"
	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cant't load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Service can't connect to the db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start the server", err)
	}

}
