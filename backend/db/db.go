package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "/data/spices.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SelectAll(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
