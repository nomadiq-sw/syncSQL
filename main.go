package main

import (
	"os"

	"github.com/nomadiq-sw/syncSQL/cmd/dbconn"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbconn.Connect_PGSQL(os.Getenv("POSTGRESQL_URL"))
	dbconn.Connect_MySQL(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_DB_NAME"),
		os.Getenv("MYSQL_HOST_PORT"),
	)
}