package database

import (
	"strings"
)

type CRUD interface {
	Get()
}

func (d *Database) Read(mask string) []Printer {
	var query string
	if mask == "" {
		query = "SELECT name, ip FROM print_printer"
	} else {
		query = "SELECT name, ip FROM print_printer WHERE name LIKE ?"
	}
	rows, _ := d.db.Query(query, strings.Replace(mask, "*", "%", -1))
	defer rows.Close()
	var printers []Printer
	for rows.Next() {
		p := &Printer{}
		rows.Scan(&p.Name, &p.IP)
		printers = append(printers, *p)
	}
	return printers
}
