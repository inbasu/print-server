package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Printer struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type DatabaseConnection interface {
	Connect()
	Close()
}

type Database struct {
	db *sql.DB
}

func (d *Database) Connect() {
	connStr := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	d.db = db
}

func (d *Database) Close() {
	d.db.Close()
}
