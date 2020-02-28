package model

import (
	"database/sql"
	"log"
)

func runQuery(query string) *sql.Rows {
	db, err := sql.Open("mysql",
		"root:senha123@tcp(127.0.0.1:3306)/gotostndb")
	if err != nil {
		log.Fatal(err)
	}
	queryResults, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return queryResults
}