package config

import (
	"crud-golang/helper"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres goalng driver
	"github.com/rs/zerolog/log"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Northern"
	dbname = "test"
)


func DatabaseConnection () *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable", host, port,  user, password, dbname )

	db, err := sql.Open("postgres", sqlInfo)

	helper.PanicIfError(err)

	log.Info().Msg("Connection establish with DB")
	return db
}