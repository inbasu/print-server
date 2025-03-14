package database

import (
	"database/sql"
	"errors"
	"strings"
)

type CRUD interface {
	Create()
	Read()
}

func (d *Database) Create(name string, ip string) error {
	var id int
	err := d.db.QueryRow("INSERT INTO printers (id, name, ip) VALUES (default, $1, $2) RETURNING id", name, ip).Scan(&id)
	if err != nil {
		/* залогировать здесь */
		return errors.New("Database error")
	}
	return nil
}

func (d *Database) Read(mask string) []Printer {
	var rows *sql.Rows
	if mask == "" {
		rows = d.getAll()
	} else {
		rows = d.getByName(strings.Replace(mask, "*", "%", -1))
	}
	defer rows.Close()
	var printers []Printer
	for rows.Next() {
		p := &Printer{}
		rows.Scan(&p.Name, &p.IP)
		printers = append(printers, *p)
	}
	return printers
}

func (d *Database) getAll() *sql.Rows {
	rows, err := d.db.Query("SELECT name, ip FROM printers")
	if err != nil {
	}
	return rows
}

func (d *Database) getByName(name string) *sql.Rows {
	rows, err := d.db.Query("SELECT name, ip FROM printers WHERE name LIKE $1", name)
	if err != nil {
	}
	return rows
}
