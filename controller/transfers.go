package controller

import (
	"encoding/json"
	"net/http"
	"go-tostn/views"
)

func Transfers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Transfers {
				Id: 1,
				AccountOriginId: 2,
				AccountDestinationId: 2,
				Amount: 14.0,
				CreatedAt: "fgfdg",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
