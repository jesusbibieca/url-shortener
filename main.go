package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jesusbibieca/url-shortener/api"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
	"github.com/jesusbibieca/url-shortener/environment"
	"github.com/jesusbibieca/url-shortener/store"
)

func main() {
	config, err := environment.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	connPool, err := pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	dbStore := db.NewStore(connPool)
	server := api.NewServer(dbStore)

	// Redis
	// move this to the api server ???
	store.InitializeStore()

	err = server.Start(config.AppAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}
