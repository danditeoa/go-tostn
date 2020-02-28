package model

import (
	structs "go-tostn/views"
)

func GetAccounts() []structs.Account {
	gotostndb := runQuery("SELECT * FROM Accounts")
	arrayAccount := []structs.Account{}
	for gotostndb.Next() {
		var acct = structs.Account{}
		gotostndb.Scan(&acct.Id, &acct.Name, &acct.Cpf, &acct.Balance, &acct.CreatedAt)
		arrayAccount = append(arrayAccount, acct)
	}
	return arrayAccount
}