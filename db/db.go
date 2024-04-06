package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/rs/zerolog/log"
)

func NewPQConn(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	initDB(db)
	return db, nil
}


func initDB(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("")
		return
	}
	log.Info().Msg("Database Initalized")
}
