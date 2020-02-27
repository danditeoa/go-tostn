package controller

import (
	"encoding/json"
	"net/http"
	"go-tostn/views"
)

func Accounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Account {
				Id:        1,
				Name:      "jsame",
				Cpf:       "jsf",
				Balance:   0,
				CreatedAt: "json:ed_at",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}