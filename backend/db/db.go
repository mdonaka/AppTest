package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DB interface {
	SelectAll() ([]Spices, error)
	Close()
}

type SqliteDB struct {
	Conn *sql.DB
}

type Spices struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
	Family string `json:"family"`
}

func NewSqliteDB(dataSourceName string) DB {
	conn, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return &SqliteDB{Conn: conn}
}

func (db *SqliteDB) Close() {
	if db.Conn != nil {
		db.Conn.Close()
	}
}

func (db *SqliteDB) SelectAll() ([]Spices, error) {
	rows, err := db.Conn.Query("SELECT * FROM spices")
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
