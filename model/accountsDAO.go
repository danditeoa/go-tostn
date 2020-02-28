package model

import (
	"database/sql"
	structs "go-tostn/views"
	"log"
)

func GetAccounts() []structs.Account {
	db, err := sql.Open("mysql",
		"root:senha123@tcp(127.0.0.1:3306)/gotostndb")
	if err != nil {
		log.Fatal(err)
	}
	gotostndb, err := db.Query("SELECT * FROM Accounts")
	arrayAccount := []structs.Account{}
	for gotostndb.Next() {
		var acct = structs.Account{}
		gotostndb.Scan(&acct.Id, &acct.Name, &acct.Cpf, &acct.Balance, &acct.CreatedAt)
		arrayAccount = append(arrayAccount, acct)
	}
	defer db.Close()
	return arrayAccount
}