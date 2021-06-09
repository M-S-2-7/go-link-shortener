package main

import (
	"database/sql"
	"log"
)

var connStr = "host=0.0.0.0 password=123456 user=postgres dbname=shortener_links sslmode=verify-full"

func SetLink(id string, link string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO links VALUES ($1, $2)`, id, link)
	if err != nil {
		log.Panic(err)
	}
}

func GetLink(id string) string{
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query(`SELECT link FROM links WHERE id = $1`, id)
	if err != nil {
		log.Panic(err)
	}

	columns, err := res.Columns()

	if err != nil {
		log.Panic(err)
	}

	return columns[0]
}
