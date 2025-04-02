package api

import (
	"backend/db"
	"encoding/json"
	"net/http"
	"strconv"
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

func (handler *HandlerWithDB) CheckSpiceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	alias := r.URL.Query().Get("alias")
	taste := r.URL.Query().Get("taste")
	flavor := r.URL.Query().Get("flavor")
	family := r.URL.Query().Get("family")
	origin := r.URL.Query().Get("origin")
	if idStr == "" || name == "" || alias == "" || taste == "" || flavor == "" || family == "" || origin == "" {
		http.Error(w, "Missing one or more parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	spice, err := handler.DB.SelectByID(id)
	if err != nil {
		http.Error(w, "Spice not found", http.StatusNotFound)
		return
	}

	match := spice.Name == name && spice.Flavor == flavor && spice.Family == family

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"match": match})
}
