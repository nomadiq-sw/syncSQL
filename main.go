package main

import (
	"database/sql"
	"os"

	"github.com/nomadiq-sw/syncSQL/cmd/dbconn"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	pg_chan := make(chan *sql.DB)
	my_chan := make(chan *sql.DB)

	go dbconn.Connect_PGSQL(os.Getenv("POSTGRESQL_URL"), pg_chan)
	go dbconn.Connect_MySQL(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_DB_NAME"),
		os.Getenv("MYSQL_HOST_PORT"),
		my_chan,
	)
	<- pg_chan
	<- my_chan
}