package dbconn

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func connect_db(driver string, url string) *sql.DB {
	db, err := sql.Open(driver, url)
	if err != nil {
		log.Fatalf("Unable to connect to %v database: %v\n", driver, err)
	}

	log.Printf("Successfully connected to %v database\n", driver)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping %v database: %v\n", driver, err)
	}

	log.Printf("Successfully pinged %v database\n", driver)

	return db
}

func Connect_PGSQL(url string, connect chan *sql.DB) {
	pg := connect_db("pgx", url)
	if connect != nil {
		connect <- pg
	}
}

func Connect_MySQL(user string, password string, db_name string, host_port string, connect chan *sql.DB) {
	cfg := mysql.Config{
		User: user,
		Passwd: password,
		Net: "tcp",
		DBName: db_name,
		Addr: host_port,
		AllowNativePasswords: true,
	}

	my := connect_db("mysql", cfg.FormatDSN())
	if connect != nil {
		connect <- my
	}
}
