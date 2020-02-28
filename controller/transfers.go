package controller

import (
	"encoding/json"
	"go-tostn/model"
	"net/http"
)

func Transfers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := model.GetTransfers()
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
