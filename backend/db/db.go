package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Spices struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
	Family string `json:"family"`
}

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "/data/spices.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SelectAll(db *sql.DB) ([]Spices, error) {
	rows, err := db.Query("SELECT * FROM spices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spices []Spices
	for rows.Next() {
		var spice Spices

		err := rows.Scan(&spice.Id, &spice.Name, &spice.Flavor, &spice.Family)
		if err != nil {
			return nil, err
		}

		spices = append(spices, spice)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return spices, nil
}
