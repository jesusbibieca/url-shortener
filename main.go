package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jesusbibieca/url-shortener/api"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
	"github.com/jesusbibieca/url-shortener/environment"
	"github.com/jesusbibieca/url-shortener/store"
	_ "github.com/lib/pq"
)

func main() {
	config, err := environment.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	dbStore := db.NewStore(conn)
	server := api.NewServer(dbStore)

	// Redis
	// move this to the api server ???
	store.InitializeStore()

	err = server.Start(config.AppAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}
