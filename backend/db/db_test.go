package db_test

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	dbpkg "backend/db"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T, data []dbpkg.User) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	assert.Nil(t, err)

	for _, entry := range data {
		_, err = db.Exec("INSERT INTO users (id, name, age) VALUES (?, ?, ?)", entry.Id, entry.Name, entry.Age)
		assert.Nil(t, err)
	}

	return db
}

// SelectAllが成功するかのテスト
func TestSelectUser_Success(t *testing.T) {
	testcases := []struct {
		name     string
		setup    []dbpkg.User
		expected []dbpkg.User
	}{
		{
			name: "select user",
			setup: []dbpkg.User{
				{1, "Alice", 20},
			},
			expected: []dbpkg.User{{1, "Alice", 20}},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db := setupTestDB(t, tt.setup)

			users, err := dbpkg.SelectAll(db)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, users)
		})
	}
}
