package model

import (
	"fmt"
	"time"
)

func CreateAccount(name, cpf, balance, datetime) error {
	insertQ, err := con.Query("INSERT INTO Accounts VALUES(?, ?, ?, ?, ?)", name, cpf, balance, Now())
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
