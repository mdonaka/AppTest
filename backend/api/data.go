package api

import (
	"backend/db"
	"encoding/json"
	"net/http"
)

type HandlerWithDB struct {
	DB db.DB
}

func (handler *HandlerWithDB) DataHandler(w http.ResponseWriter, r *http.Request) {
	response, err := handler.DB.SelectAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
