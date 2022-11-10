package db

import (
	"os"
        "log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func MySQLConnection() *sql.DB {

	configDB := mysql.Config{
		User:   os.Getenv("USERNAME"),
		Passwd: os.Getenv("PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DATABASE"),
	}

        db, err := sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

