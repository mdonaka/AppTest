package api

import (
	dbpkg "backend/db"
	"encoding/json"
	"net/http"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	mydb := dbpkg.Open()
	defer mydb.Close()
	response, err := dbpkg.SelectAll(mydb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
