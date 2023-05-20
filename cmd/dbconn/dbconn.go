package dbconn

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func connect_db(driver string, url string) {
	db, err := sql.Open(driver, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to %v database: %v\n", driver, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Successfully connected to %v database!\n", driver)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping %v database: %v\n", driver, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Successfully pinged %v database!\n", driver)
}

func Connect_PGSQL(url string) {
	connect_db("pgx", url)
}

func Connect_MySQL(user string, password string, db_name string, host_port string) {
	cfg := mysql.Config{
		User: user,
		Passwd: password,
		Net: "tcp",
		DBName: db_name,
		Addr: host_port,
		AllowNativePasswords: true,
	}

	connect_db("mysql", cfg.FormatDSN())
}
