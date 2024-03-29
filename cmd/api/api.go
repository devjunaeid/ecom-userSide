package api

import (
	"database/sql"
	"net/http"

	"github.com/devjunaeid/ecom-userSide/service/user"
	log "github.com/rs/zerolog/log"
)

type apiServer struct {
	address  string
	database *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *apiServer {
	return &apiServer{
		address:  addr,
		database: db,
	}
}
func (s *apiServer) Run() error {

	router := http.NewServeMux()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(router)

	log.Info().Msgf("Server Running on port, %s", s.address)

	return http.ListenAndServe(s.address, router)
}
