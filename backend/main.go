package main

import (
	"backend/api"
	"backend/db"
	"fmt"
	"net/http"
)

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpoint := r.URL.Path
		fmt.Println("endpoint:", endpoint)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next(w, r)
	}
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func main() {
	database := db.NewSqliteDB("/data/spices.db")
	defer database.Close()
	dataHandlerStruct := &api.HandlerWithDB{DB: database}

	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware(notfoundHandler))
	mux.HandleFunc("/data", middleware(dataHandlerStruct.DataHandler))
	mux.HandleFunc("/check", middleware(dataHandlerStruct.CheckSpiceHandler))

	port := ":8080"
	fmt.Printf("Server started at %s\n", port)
	http.ListenAndServe(port, mux)
}
