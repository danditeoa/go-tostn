package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/accounts", Accounts())
	mux.HandleFunc("/transfers", Transfers())
	return mux
}