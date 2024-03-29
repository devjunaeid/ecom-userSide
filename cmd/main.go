package main

import (
	"github.com/devjunaeid/ecom-userSide/cmd/api"
	log "github.com/rs/zerolog/log"
)

func main() {
	server := api.NewApiServer(":8080", nil)
	var err error
	err = server.Run()

	if err != nil {
		log.Fatal().Err(err).Msg("")
		return
	}

}
