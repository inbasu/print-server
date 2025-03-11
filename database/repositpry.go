package database

import (
	"database/sql"
	"regexp"
	"strings"
)

type CRUD interface {
	Create()
	Read()
}

func (d *Database) Create(name string, ip string) *Printer {
	if !isValidIp(ip) {
		return nil
	}
	p := &Printer{}
	err := d.db.QueryRow("INSERT INTO printers (id, name, ip) VALUES (default, $1, $2)", name, ip).Scan(&p.Name, &p.Name)
	if err != nil {
	}
	return p

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

func isValidIp(ip string) bool {
	ipPattern, _ := regexp.Compile("^((?:[01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])$")
	return ipPattern.Match([]byte(ip))
}
