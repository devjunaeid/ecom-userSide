package main

import (
	"github.com/devjunaeid/ecom-userSide/cmd/api"
	"github.com/devjunaeid/ecom-userSide/config"
	"github.com/devjunaeid/ecom-userSide/db"
	log "github.com/rs/zerolog/log"
)

func main() {
	// Initializing new Db connection.
	// Returns db and an erros. takse Database connection url.
	db, err := db.NewPQConn(config.Envs.DbConnectionUrl) //load from .env file via config module.
	//Checking for database connection erros.
	if err != nil {
		log.Error().Msg("Failed to Connect To the Database!!")
	} else {
		log.Info().Msg("Connected to the database.")
	}

	server := api.NewApiServer(":8080", db)
	err = server.Run()

	if err != nil {
		log.Fatal().Err(err).Msg("")
		return
	}

}
