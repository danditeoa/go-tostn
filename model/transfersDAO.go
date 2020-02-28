package model

import (
	structs "go-tostn/views"
)

func GetTransfers() []structs.Transfers {
	gotostndb := runQuery("SELECT * FROM Transfers")
	var listTransfers []structs.Transfers
	for gotostndb.Next() {
		var transfers = structs.Transfers{}
		gotostndb.Scan(&transfers.Id, &transfers.AccountOriginId, &transfers.AccountDestinationId, &transfers.Amount, &transfers.CreatedAt)
		listTransfers = append(listTransfers, transfers)
	}
	return listTransfers
}
