package database

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
)

type Printer struct {
	Name string
	IP   string
}

type DatabaseConnection interface {
	Connect()
	Close()
}

type Database struct {
	db *sql.DB
}

func (d *Database) Connect() {
	db, _ := sql.Open("sqlite", "./db.sqlite3")
	d.db = db
}

func (d *Database) Close() {
	d.db.Close()
}
