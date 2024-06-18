package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mogmoggy/garage-booking-backend/api"
	db "github.com/mogmoggy/garage-booking-backend/db/sqlc"
	"github.com/mogmoggy/garage-booking-backend/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store, *config)

	if err = server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
