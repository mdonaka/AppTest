package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Select_all() []User {
	db, err := sql.Open("sqlite3", "/data/spices.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}
