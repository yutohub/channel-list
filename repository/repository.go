package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "channel_list",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	// Actual connection test
	pingErr := db.Ping()
	if pingErr != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Println("Connected!")
}
