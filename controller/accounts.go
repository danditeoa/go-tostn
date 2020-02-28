package controller

import (
	"encoding/json"
	"go-tostn/model"
	"net/http"
)

func Accounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := model.GetAccounts()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}