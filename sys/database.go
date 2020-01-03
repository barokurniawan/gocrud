package sys

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connection *sql.DB
	host       string
	username   string
	dbname     string
	password   string
}

func (db *Database) CreateConnection() *Database {
	db.dbname = "gocrud_db"
	db.host = "127.0.0.1"
	db.username = "root"
	db.password = "@@@udanup"

	var conn *sql.DB
	var err error

	conn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3309)/%s", db.username, db.password, db.host, db.dbname))
	if err != nil {
		panic(err)
	}

	db.Connection = conn
	return db
}

func (db *Database) GetCurrentConnection() *Database {
	return db
}
