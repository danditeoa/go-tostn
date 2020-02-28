package model

import (
	"database/sql"
	structs "go-tostn/views"
	"log"
)

func GetTransfers() []structs.Transfers {
	db, err := sql.Open("mysql",
		"root:senha123@tcp(127.0.0.1:3306)/gotostndb")
	if err != nil {
		log.Fatal(err)
	}
	gotostndb, err := db.Query("SELECT * FROM Transfers")
	listTransfers := []structs.Transfers{}
	for gotostndb.Next() {
		var transfers = structs.Transfers{}
		gotostndb.Scan(&transfers.Id, &transfers.AccountOriginId, &transfers.AccountDestinationId, &transfers.Amount, &transfers.CreatedAt)
		listTransfers = append(listTransfers, transfers)
	}
	defer db.Close()
	return listTransfers
}
