package api

import (
	"encoding/json"
	"net/http"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	response := map[string]any{
		"message": "Hello, World!",
		"status":  "success",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
